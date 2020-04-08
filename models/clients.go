package models

import (
    "errors"
    "github.com/jinzhu/gorm"
	"github.com/Htech/Htech-api/db"
	"github.com/Htech/Htech-api/httpmodels"
)

type Clients struct {
    gorm.Model
	Name                string
	Document            string          `gorm:"unique_index:idx_document_type_document"`
    TypeDocuments       TypeDocuments   `gorm:"foreignkey:TypeDocumentsID"`
    TypeDocumentsID     uint            `gorm:"unique_index:idx_document_type_document"`
	Email               string          `gorm:"unique_index:idx_email"`
    Phone               string          `gorm:"size:13"`
	Status              bool
    Users               Users           `gorm:"foreignkey:UsersID"`
    UsersID             uint
}

func GetClientByEmail(email string) *Clients {
	client := &Clients{}
	err := db.DB.Where("email = ?", email).First(client).Error
	if err != nil {
		return nil
	}
	return client
}

func GetClientByDocumentAndType(document string, TypeDocumentsID uint) *Clients {
	client := &Clients{}
	err := db.DB.Where("document = ? AND type_documents_id = ?", document, TypeDocumentsID).Find(client).Error
	if err != nil {
		return nil
	}
	return client
}

func GetClientsByDocumentAndType(document string, TypeDocumentsID uint) []Clients {
	client := []Clients{}
	err := db.DB.Where("document = ? AND type_documents_id = ?", document, TypeDocumentsID).Find(&client).Error
	if err != nil {
		return nil
	}
	return client
}

func GetClients() []Clients {
	client := []Clients{}
	err := db.DB.Find(&client).Error
	if err != nil {
		return nil
	}
	return client
}

func CreateClients(Name string, Document string, TypeDocumentsID uint, Email string, Phone string, Status bool, UsersID uint) error {
    client := &Clients{
        Name:               Name,
        Document:           Document,
        TypeDocumentsID:    TypeDocumentsID,
        Email:              Email,
        Phone:              Phone,
        Status:             true,
        UsersID:            UsersID,
    }
    err := db.DB.Create(&client).Error
    if err != nil {
        return err
    } else {
        return nil
    }
}

func UpdateClients(body *httpmodels.UpdateClientRequest) error {
	client := &Clients{}
	err := db.DB.Where("id = ?", body.ID).First(client).Error
    if err != nil {
        return err
    }
    validEmail := GetClientByEmail(*body.Email)
    if validEmail != nil && validEmail.ID != client.ID{
        return errors.New("El email ya existe")
    }
    if body.Name != nil && *body.Name != client.Name{
        client.Name = *body.Name
    }
    if body.Email != nil && *body.Email != client.Email{
        client.Email = *body.Email
    }
    if body.Phone != nil && *body.Phone != client.Phone{
        client.Phone = *body.Phone
    }
    err = client.Save()
    if err != nil {
        return err
    } else {
        return nil
    }
}

func (c *Clients) Save() error {
	return db.DB.Save(c).Error
}
