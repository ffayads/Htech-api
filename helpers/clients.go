package helpers

import (
    "fmt"
    "errors"
	"github.com/Htech/Htech-api/models"
	"github.com/Htech/Htech-api/httpmodels"
)

func CreateClient(body *httpmodels.CreateClientRequest, email string) (error, string) {
    typeDocument := models.GetTypeDocumentByCode(body.TypeDocument)
    if typeDocument == nil {
        return nil,"El tipo de documento no existe"
    }
    err := models.GetClientByDocumentAndType(body.Document, typeDocument.ID)
    if err != nil {
        return errors.New("El documento ya existe"),"El documento ya existe"
    }
    user := models.GetUserByEmail(email)
    if user == nil {
        return errors.New("El email no existe"),"El email no existe"
    }
    errCreate := models.CreateClients(body.Name, body.Document, typeDocument.ID, body.Email, body.Phone, true, user.ID)
    if errCreate != nil{
        return errCreate,"No se pudo crear el cliente"
    }
    return nil,"El cliente fue creado con exito"
}

func GetClients(body *httpmodels.GetClientsRequest) (*[]httpmodels.GetClientResponse, string) {
    modelClients := []models.Clients{}
    fmt.Println("Document: ",body.Document)
    fmt.Println("Type: ",body.TypeDocument)
    if body.Document == nil || body.TypeDocument == nil{
        modelClients = models.GetClients()
    } else {
        typeDocument := models.GetTypeDocumentByCode(*body.TypeDocument)
        if typeDocument == nil {
            return nil,"El tipo de documento no existe"
        }
        modelClients = models.GetClientsByDocumentAndType(*body.Document, typeDocument.ID)
    }
    clients := []httpmodels.GetClientResponse{}
    for _, element := range modelClients {
        typeDocument := models.GetTypeDocumentByID(element.TypeDocumentsID)
        client := httpmodels.GetClientResponse{
            ID:                 element.ID,
            Name:               element.Name,
            Document:           element.Document,
            TypeDocuments:      typeDocument.Description,
            Email:              element.Email,
            Phone:              element.Phone,
            Status:             element.Status,
        }
        clients = append(clients, client)
    }
    if clients == nil{
        return nil,"No se encontro el cliente"
    }
    return &clients,""
}

func UpdateClient(body *httpmodels.UpdateClientRequest) (error, string) {
    errCreate := models.UpdateClients(body)
    if errCreate != nil{
        return errCreate,"No se pudo editar el cliente"
    }
    return nil,"El cliente fue editado con exito"
}
