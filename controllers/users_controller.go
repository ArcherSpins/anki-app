package controllers

import (
	"anki-project/models"
	"anki-project/services"
	"anki-project/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func getResponse(user models.User) (models.LoginResponse, error) {
	token, err := utils.GenerateJWT(user.ID, user.Username)

	if err != nil {
		return models.LoginResponse{}, err
	}

	response := models.LoginResponse{
		User: models.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
		Token: token,
	}

	return response, nil
}

func (ctrl *UserController) Login(c *gin.Context) {
	var loginData models.LoginData

	if err := c.ShouldBindBodyWithJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.UserService.Login(loginData.Login, loginData.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	responseUser := models.User{
		Email:    user.Email,
		ID:       user.ID,
		Username: user.Username,
		Password: user.PasswordHash,
	}

	response, err := getResponse(responseUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *UserController) Register(c *gin.Context) {
	var registerData models.User

	if err := c.ShouldBindBodyWithJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser := models.DBUser{
		Email:        registerData.Email,
		Username:     registerData.Username,
		PasswordHash: string(passwordHash),
	}

	user, err := ctrl.UserService.Register(newUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseUser := models.User{
		Email:    user.Email,
		Username: user.Username,
		Password: user.PasswordHash,
		ID:       user.ID,
	}

	response, err := getResponse(responseUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (ctrl *UserController) Edit(c *gin.Context) {
	var editUser models.EditUser

	if err := c.ShouldBindBodyWithJSON(&editUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.Abort()
		return
	}

	claims, err := utils.VerifyJWT(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.Abort()
		return
	}

	card, err := ctrl.UserService.Edit(claims.ID, editUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, card)
}
