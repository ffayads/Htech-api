package httpmodels

import (
)

type Respose struct {
  Status    bool
  Data      string
  Msg       string
  Error     error
}

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type CreateClientRequest struct {
	Name                string
	Document            string
    TypeDocument        string  `form:"type_document" json:"type_document" binding:"required"`
	Email               string
    Phone               string
}

type GetClientsRequest struct {
	Document            *string
    TypeDocument        *string
}

type GetClientResponse struct {
	ID                  uint
	Name                string
	Document            string
    TypeDocuments       string
	Email               string
    Phone               string
	Status              bool
}

type UpdateClientRequest struct {
	ID                  uint
	Name                *string
	Email               *string
    Phone               *string
}

type CreateCardRequest struct {
	Number      string
	Expiry      string
	Cvv         string
	ClientsID   uint        `form:"clients_id" json:"clients_id" binding:"required"`
}

type GetCardResponse struct {
	Number      string
	Expiry      string
}

type UpdateCardRequest struct {
	ID          uint
	Number      *string
	Expiry      *string
	Cvv         *string
}
