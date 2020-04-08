package models

import (
    "github.com/jinzhu/gorm"
	"github.com/Htech/Htech-api/db"
)

type Users struct {
    gorm.Model
	Name                string
	Document            string          `gorm:"unique_index:idx_document_type_document"`
    TypeDocuments       TypeDocuments   `gorm:"foreignkey:TypeDocumentsID"`
    TypeDocumentsID     uint            `gorm:"unique_index:idx_document_type_document"`
	Email               string          `gorm:"unique_index:idx_email"`
	Password            string
    Phone               string          `gorm:"size:13"`
	Status              bool
}

func GetUserByID(id uint) *Users {
	user := &Users{}
	err := db.DB.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil
	}
	return user
}

func GetUserByEmail(email string) *Users {
	user := &Users{}
	err := db.DB.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil
	}
	return user
}

func GetUserByDocumentAndType(document string, typeDocumentID uint) *Users {
	user := &Users{}
	err := db.DB.Where("document = ? AND type_document_id = ?", document, typeDocumentID).First(user).Error
	if err != nil {
		return nil
	}
	return user
}

