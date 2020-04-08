package helpers

import (
    "strconv"
	"github.com/Htech/Htech-api/models"
	"github.com/Htech/Htech-api/httpmodels"
)

func CreateCards(body *httpmodels.CreateCardRequest) (error, string) {
    errCreate := models.CreateCards(body.Number, body.Expiry, body.Cvv, body.ClientsID)
    if errCreate != nil{
        return errCreate,"No se pudo crear la tarjeta"
    }
    return nil,"La tarjeta fue creada con exito"
}

func GetCards(clientParam string, cardParam string) (*[]httpmodels.GetCardResponse, string) {
    modelCards := []models.Cards{}
    if cardParam != "" {
        card ,_ := strconv.Atoi(cardParam)
        cardId := uint(card)
        modelCards = models.GetCardsById(cardId)
    } else if clientParam != "" {
        client ,_ := strconv.Atoi(clientParam)
        clientId := uint(client)
        modelCards = models.GetCardsByClientId(clientId)
    } else {
        return nil,"No se encontraron tarjetas"
    }
    cards := []httpmodels.GetCardResponse{}
    for _, element := range modelCards {
        card := httpmodels.GetCardResponse{
            Number:     element.Number,
            Expiry:     element.Expiry,
        }
        cards = append(cards, card)
    }
    if cards == nil{
        return nil,"No se encontraron tarjetas"
    }
    return &cards,""
}

func UpdateCards(body *httpmodels.UpdateCardRequest) (error, string) {
    errCreate := models.UpdateCards(body)
    if errCreate != nil{
        return errCreate,"No se pudo editar la tarjeta"
    }
    return nil,"La tarjeta fue editada con exito"
}
