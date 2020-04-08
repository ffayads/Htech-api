package controllers

import (
    "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/Htech/Htech-api/helpers"
	"github.com/Htech/Htech-api/httpmodels"
)

func CreateClient(c *gin.Context){
    response := &httpmodels.Respose{}
    claims := jwt.ExtractClaims(c)
    body := &httpmodels.CreateClientRequest{}
    if err := c.ShouldBindJSON(&body); err != nil {
        response.Status = false
        response.Error = err
        c.AbortWithStatusJSON(200, response)
        return
    }
    err, errMsg := helpers.CreateClient(body, claims["email"].(string))
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

func GetClients(c *gin.Context){
    document := c.PostForm("document")
    typeDocument := c.PostForm("type_document")
    body := &httpmodels.GetClientsRequest{}
    if document == "" || typeDocument == "" {
        body = &httpmodels.GetClientsRequest{
            Document:       nil,
            TypeDocument:   nil,
        }
    } else {
        body = &httpmodels.GetClientsRequest{
            Document:       &document,
            TypeDocument:   &typeDocument,
        }
    }
    clients, errMsg := helpers.GetClients(body)
    if clients == nil {
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
                "clients":  clients,
			},
		})
    }
    return
}

func UpdateClient(c *gin.Context){
    response := &httpmodels.Respose{}
    body := &httpmodels.UpdateClientRequest{}
    if err := c.ShouldBindJSON(&body); err != nil {
        response.Status = false
        response.Error = err
        c.AbortWithStatusJSON(200, response)
        return
    }
    err, errMsg := helpers.UpdateClient(body)
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
