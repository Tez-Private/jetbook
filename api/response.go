package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int
	Data   interface{}
}

func RespondJSON(ctx *gin.Context, status int, paylaod interface{}) {
	log.Println("status ", status)
	var response ResponseData

	response.Status = status
	response.Data = paylaod

	ctx.JSON(status, response)
}

func RespondLogin(ctx *gin.Context, status int, paylaod interface{}) {
	log.Println("status ", status)
	var response ResponseData

	response.Status = status
	response.Data = paylaod

	ctx.JSON(status, response)
}
