package handler

import (
	"fmt"
	"net/http"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Reza Alfara",
		"bio":  "A Software Engineer",
	})
}

func HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "Belajar golang bareng Reza Alfara",
	})
}

func BooksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")

	ctx.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	price := ctx.Query("price")

	ctx.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

func PostBooksHandler(ctx *gin.Context) {
	var bookInput book.BookInput

	err := ctx.ShouldBindJSON(&bookInput)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			// Handle validation errors
			var errorMessages []string
			for _, e := range validationErrs {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return
		}

		// Handle other errors
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})

}
