package services

import (
	"anki-project/models"
	"anki-project/repository"
)

type CardService interface {
	CreateCard(card models.Card) (models.Card, error)
	EditCard(cardId int, card models.EditCard) (models.Card, error)
	GetCard(cardId int64) (models.Card, error)
	GetListOfCards(account_id int) ([]models.Card, error)
	DeleteCard(cardId int64) error
}

type cardService struct {
	repo repository.CardRepository
}

func NewCardService(repo repository.CardRepository) CardService {
	return &cardService{repo: repo}
}

func (c *cardService) CreateCard(card models.Card) (models.Card, error) {
	return c.repo.CreateCard(card)
}

func (c *cardService) DeleteCard(cardId int64) error {
	return c.repo.DeleteCard(cardId)
}

func (c *cardService) EditCard(cardId int, card models.EditCard) (models.Card, error) {
	return c.repo.EditCard(cardId, card)
}

func (c *cardService) GetCard(cardId int64) (models.Card, error) {
	return c.repo.GetCard(cardId)
}

func (c *cardService) GetListOfCards(account_id int) ([]models.Card, error) {
	return c.repo.GetListOfCards(account_id)
}
