package v1

import (
	"log"
	"net/http"

	api "github.com/Tez-Private/jetbook/api"
	"github.com/Tez-Private/jetbook/service/rent"
	"github.com/gin-gonic/gin"
)

//Create POST api/v1/rents
func Create(ctx *gin.Context) {
	var params rent.CreateParams
	err := ctx.BindJSON(&params)
	if err != nil {
		log.Println(err)
		api.RespondJSON(ctx, http.StatusNotFound, params)
		return
	}

	result, err := rent.Create(ctx, &params)
	if err != nil {
		log.Println(err)
		api.RespondJSON(ctx, http.StatusNotFound, result)
		return
	}
	api.RespondJSON(ctx, http.StatusCreated, result)
}

//GetAll GET api/v1/rents
func GetAll(ctx *gin.Context) {
	result, err := rent.GetAll(ctx)

	if err != nil {
		api.RespondJSON(ctx, http.StatusNotFound, result)
	}
	api.RespondJSON(ctx, http.StatusOK, result)
}

//Get GET api/v1/rents/[id]
func Get(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	result, err := rent.Get(id)
	if err != nil {
		api.RespondJSON(ctx, http.StatusNotFound, result)
		return
	}
	api.RespondJSON(ctx, http.StatusOK, result)
}

//Return PATCH api/v1/rents/[id]
func Return(ctx *gin.Context) {
	var params rent.UpdateParams
	id := ctx.Params.ByName("id")

	err := ctx.BindJSON(&params)
	if err != nil {
		api.RespondJSON(ctx, http.StatusNotFound, err)
		return
	}

	result, err := rent.Update(ctx, &params, id)
	if err != nil {
		api.RespondJSON(ctx, http.StatusNotFound, result)
		return
	}
	api.RespondJSON(ctx, http.StatusOK, result)
}

//Update PATCH api/v1/rents/[id]
func Update(ctx *gin.Context) {
	var params rent.UpdateParams
	id := ctx.Params.ByName("id")
	err := ctx.BindJSON(&params)
	if err != nil {
		api.RespondJSON(ctx, http.StatusNotFound, err)
		return
	}

	result, err := rent.Update(ctx, &params, id)
	if err != nil {
		api.RespondJSON(ctx, http.StatusNotFound, result)
		return
	}
	api.RespondJSON(ctx, http.StatusOK, result)
}

/*
//Delete DELETE api/v1/rents/[id]
func Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	err := rent.Delete(ctx, id)
	if err != nil {
		log.Println(err)
		api.RespondJSON(ctx, http.StatusNotFound, "Denied")
	}
	api.RespondJSON(ctx, http.StatusNoContent, "Not Content")
}
*/
