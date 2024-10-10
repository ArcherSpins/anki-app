package repository

import (
	"anki-project/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user models.DBUser) (models.DBUser, error)
	Login(login string, password string) (models.DBUser, error)
	Edit(userId int, user models.EditUser) (models.User, error)
	GetAllUsers(page int) ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Register(user models.DBUser) (models.DBUser, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) Login(login string, password string) (models.DBUser, error) {
	var user models.DBUser
	err := r.db.Where("email = ?", login).Or("username = ?", login).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.DBUser{}, errors.New("user not found")
		}
		return models.DBUser{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return models.DBUser{}, errors.New("incorrect password")
	}

	return user, nil
}

func (r *userRepository) Edit(userId int, updatedUser models.EditUser) (models.User, error) {
	var user models.EditUser

	err := r.db.Where("id = ?", userId).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}

	if updatedUser.Email != nil {
		user.Email = updatedUser.Email
	}

	if updatedUser.Username != nil {
		user.Username = updatedUser.Username
	}

	if updatedUser.PasswordHash != nil {
		user.PasswordHash = updatedUser.PasswordHash
	}

	err = r.db.Save(&user).Error

	if err != nil {
		return models.User{}, err
	}

	var newUser models.User

	err = r.db.Where("id = ?", userId).First(&newUser).Error

	if err != nil {
		return models.User{}, err
	}

	return newUser, nil
}

func (r *userRepository) GetAllUsers(page int) ([]models.User, error) {

	// users, err := r.db.

	return []models.User{}, nil
}
