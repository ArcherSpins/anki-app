package controllers

import (
	"anki-project/models"
	"anki-project/services"
	"anki-project/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CardsController struct {
	cardService services.CardService
}

func NewCardsController(cardService services.CardService) *CardsController {
	return &CardsController{cardService: cardService}
}

func (ctrl *CardsController) CreateCard(c *gin.Context) {
	var card models.Card

	if err := c.ShouldBindBodyWithJSON(&card); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Card Data: %+v\n", card)

	card, err := ctrl.cardService.CreateCard(card)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, card)
}

func (ctrl *CardsController) DeleteCard(c *gin.Context) {
	cardIdStr := c.Param("id")

	cardId, err := strconv.ParseInt(cardIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
		return
	}

	err = ctrl.cardService.DeleteCard(cardId)

	if err != nil {
		if err.Error() == "card not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found card with this id"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		return
	}

	c.Status(http.StatusNoContent)
}

func (ctrl *CardsController) EditCard(c *gin.Context) {
	var updatedCard models.EditCard

	if err := c.ShouldBindJSON(&updatedCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	card, err := ctrl.cardService.EditCard(updatedCard.CardId, updatedCard)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, card)
}

func (ctrl *CardsController) GetCard(c *gin.Context) {
	cardIdStr := c.Param("id")

	cardId, err := strconv.ParseInt(cardIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
		return
	}

	card, err := ctrl.cardService.GetCard(cardId)

	if err != nil {
		if err.Error() == "card not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found card with this id"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, card)
}

func (ctrl *CardsController) GetListOfCards(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		c.Abort()
		return
	}

	// Ожидаемый формат заголовка: "Bearer <token>"
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

	account_id := claims.ID

	fmt.Println(claims, account_id)

	cards, err := ctrl.cardService.GetListOfCards(account_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cards)
}
