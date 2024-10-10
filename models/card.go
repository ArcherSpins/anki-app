package models

import "time"

type Card struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	AccountID  int       `json:"account_id"`
	FrontLabel *string   `json:"front_label"`
	BackLabel  *string   `json:"back_label"`
	Visible    *bool     `json:"visible"`
	Color      *string   `json:"color"`
	CreatedAt  time.Time `json:"created_at"`
}

type EditCard struct {
	CardId     int       `json:"card_id"`
	FrontLabel *string   `json:"front_label"`
	BackLabel  *string   `json:"back_label"`
	Visible    *bool     `json:"visible"`
	Color      *string   `json:"color"`
	CreatedAt  time.Time `json:"created_at"`
}
