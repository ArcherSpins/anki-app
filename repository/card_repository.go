package repository

import (
	"anki-project/models"
	"errors"

	"gorm.io/gorm"
)

type CardRepository interface {
	CreateCard(card models.Card) (models.Card, error)
	EditCard(cardId int, card models.EditCard) (models.Card, error)
	GetCard(cardId int64) (models.Card, error)
	GetListOfCards(account_id int) ([]models.Card, error)
	DeleteCard(cardId int64) error
}

type cardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &cardRepository{db: db}
}

func (r *cardRepository) DeleteCard(cardId int64) error {
	err := r.db.Where("id = ?", cardId).Delete(&models.Card{}).Error
	return err
}

func (r *cardRepository) EditCard(cardId int, card models.EditCard) (models.Card, error) {
	var oldCard models.Card
	err := r.db.Where("id = ?", cardId).First(&oldCard).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Card{}, errors.New("not found card with this id")
		}
		return models.Card{}, err
	}

	if card.Visible != nil {
		oldCard.Visible = card.Visible
	}

	if card.FrontLabel != nil {
		oldCard.FrontLabel = card.FrontLabel
	}

	if card.BackLabel != nil {
		oldCard.BackLabel = card.BackLabel
	}

	if card.Color != nil {
		oldCard.Color = card.Color
	}

	err = r.db.Save(&oldCard).Error

	if err != nil {
		return models.Card{}, err
	}

	return oldCard, nil
}

func (r *cardRepository) GetCard(cardId int64) (models.Card, error) {
	var card models.Card
	err := r.db.Where("id = ?", cardId).First(&card).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Card{}, errors.New("card not found")
		}
		return models.Card{}, err
	}

	return card, nil
}

func (r *cardRepository) GetListOfCards(account_id int) ([]models.Card, error) {
	var cards []models.Card
	err := r.db.Where("account_id = ?", account_id).Find(&cards).Error

	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (r *cardRepository) CreateCard(card models.Card) (models.Card, error) {
	err := r.db.Create(&card).Error
	return card, err
}
