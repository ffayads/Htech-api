package models

import (
    "github.com/jinzhu/gorm"
	"github.com/Htech/Htech-api/db"
	"github.com/Htech/Htech-api/httpmodels"
)

type Cards struct {
    gorm.Model
	Number      string  `gorm:"unique_index:idx_number_clients_id"`
	Expiry      string
	Cvv         string
	Clients     Clients
	ClientsID     uint    `gorm:"unique_index:idx_number_clients_id"`
}

func GetCardsById(id uint) []Cards {
	cards := []Cards{}
	err := db.DB.Where("id = ?", id).Find(&cards).Error
	if err != nil {
		return nil
	}
	return cards
}

func GetCardsByClientId(clientId uint) []Cards {
	card := []Cards{}
	err := db.DB.Where("clients_id = ?", clientId).Find(&card).Error
	if err != nil {
		return nil
	}
	return card
}

func CreateCards(number string, expiry string, cvv string, clientsID uint) error {
    cards := &Cards{
        Number:     number,
        Expiry:     expiry,
        Cvv:        cvv,
        ClientsID: clientsID,
    }
    err := db.DB.Create(&cards).Error
    if err != nil {
        return err
    } else {
        return nil
    }
}

func UpdateCards(body *httpmodels.UpdateCardRequest) error {
	card := &Cards{}
	err := db.DB.Where("id = ?", body.ID).First(card).Error
    if err != nil {
        return err
    }
    if body.Number != nil && *body.Number != card.Number{
        card.Number = *body.Number
    }
    if body.Cvv != nil && *body.Cvv != card.Cvv{
        card.Cvv = *body.Cvv
    }
    if body.Expiry != nil && *body.Expiry != card.Expiry{
        card.Expiry = *body.Expiry
    }
    err = card.Save()
    if err != nil {
        return err
    } else {
        return nil
    }
}

func (c *Cards) Save() error {
	return db.DB.Save(c).Error
}
