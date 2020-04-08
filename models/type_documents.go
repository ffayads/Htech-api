package models

import (
    "github.com/jinzhu/gorm"
	"github.com/Htech/Htech-api/db"
)

type TypeDocuments struct {
    gorm.Model
	Code            string
	Description     string
}

func GetTypeDocumentByID(id uint) *TypeDocuments {
	typeDocuments := &TypeDocuments{}
	err := db.DB.Where("id = ?", id).First(typeDocuments).Error
	if err != nil {
		return nil
	}
	return typeDocuments
}

func GetTypeDocumentByCode(code string) *TypeDocuments {
	typeDocuments := &TypeDocuments{}
	err := db.DB.Where("code = ?", code).First(typeDocuments).Error
	if err != nil {
		return nil
	}
	return typeDocuments
}

func CreateTypeDocuments(Code string, Description string) error {
    typeDocuments := &TypeDocuments{
        Code:           Code,
        Description:    Description,
    }
    err := db.DB.Create(&typeDocuments).Error
    if err != nil {
        return err
    } else {
        return nil
    }
}
