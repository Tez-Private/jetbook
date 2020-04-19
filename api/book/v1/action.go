package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	api "github.com/Tez-Private/jetbook/api"
	"github.com/Tez-Private/jetbook/service/book"
	"github.com/gin-gonic/gin"
)

//Create POST api/v1/books
func Create(ctx *gin.Context) {
	var getIsbn book.GetIsbn
	err := ctx.BindJSON(&getIsbn)
	if err != nil {
		log.Println(err)
		api.RespondJSON(ctx, http.StatusNotFound, getIsbn)
		return
	}

	// Get Book Info
	baseurl := "https://www.googleapis.com/books/v1/volumes?q=isbn:"
	url := baseurl + getIsbn.Isbn

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	jsonBytes := ([]byte)(byteArray)
	data := new(book.GoogleBooksAPI)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	result, err := book.Create(ctx, data)
	if err != nil {
		log.Println(err)
		api.RespondJSON(ctx, http.StatusNotFound, result)
		return
	}
	api.RespondJSON(ctx, http.StatusCreated, result)
}

//GetAll GET api/v1/books
func GetAll(ctx *gin.Context) {
	result, err := book.GetAll(ctx)

	if err != nil {
		api.RespondJSON(ctx, http.StatusNotFound, result)
	}
	api.RespondJSON(ctx, http.StatusOK, result)
}

//Get GET api/v1/books/[isbn]
func Get(ctx *gin.Context) {
	isbn := ctx.Params.ByName("isbn")
	result, err := book.Get(isbn)
	if err != nil {
		api.RespondJSON(ctx, http.StatusNotFound, result)
		return
	}
	api.RespondJSON(ctx, http.StatusOK, result)
}

//Delete DELETE api/v1/books/[isbn]
func Delete(ctx *gin.Context) {
	isbn := ctx.Params.ByName("isbn")

	err := book.Delete(ctx, isbn)
	if err != nil {
		log.Println(err)
		api.RespondJSON(ctx, http.StatusNotFound, "Denied")
	}
	api.RespondJSON(ctx, http.StatusNoContent, "Not Content")
}
