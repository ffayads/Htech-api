package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/Htech/Htech-api/helpers"
	"github.com/Htech/Htech-api/httpmodels"
)

func CreateCards(c *gin.Context){
    response := &httpmodels.Respose{}
    body := &httpmodels.CreateCardRequest{}
    if err := c.ShouldBindJSON(&body); err != nil {
        response.Status = false
        response.Error = err
        c.AbortWithStatusJSON(200, response)
        return
    }
    err, errMsg := helpers.CreateCards(body)
    if err != nil {
        response.Status = false
        response.Error = err
        response.Msg = errMsg
        c.AbortWithStatusJSON(200, response)
    } else {
        response.Status = true
        response.Msg = errMsg
		c.JSON(200, response)
    }
    return
}

func GetCards(c *gin.Context){
    clientId := c.PostForm("clients_id")
    cardId := c.PostForm("card_id")
    cards, errMsg := helpers.GetCards(clientId, cardId)
    if cards == nil {
        c.JSON(200, gin.H{
			"status": false,
			"err":    "",
			"msg":    errMsg,
			"data": gin.H{
			},
		})
    } else {
        c.JSON(200, gin.H{
			"status": true,
			"err":    "",
			"msg":    errMsg,
			"data": gin.H{
                "cards":  cards,
			},
		})
    }
    return
}

func UpdateCards(c *gin.Context){
    response := &httpmodels.Respose{}
    body := &httpmodels.UpdateCardRequest{}
    if err := c.ShouldBindJSON(&body); err != nil {
        response.Status = false
        response.Error = err
        c.AbortWithStatusJSON(200, response)
        return
    }
    err, errMsg := helpers.UpdateCards(body)
    if err != nil {
        response.Status = false
        response.Error = err
        response.Msg = errMsg
        c.AbortWithStatusJSON(200, response)
    } else {
        response.Status = true
        response.Msg = errMsg
		c.JSON(200, response)
    }
    return
}
