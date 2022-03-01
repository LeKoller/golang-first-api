package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lekoller/api-with-go/database"
	"github.com/lekoller/api-with-go/models"
)

func RetrieveBook(ctx *gin.Context) {
	param_id := ctx.Param("id")
	id, err := strconv.Atoi(param_id)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "id has to be integer",
		})

		return
	}

	db := database.GetDatabase()
	var book models.Book

	err = db.First(&book, id).Error

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "book not found: " + err.Error(),
		})

		return
	}

	ctx.JSON(200, book)
}

func CreateBook(ctx *gin.Context) {
	var book models.Book
	err := ctx.ShouldBindJSON(&book)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot bind JSON to a book object",
		})

		return
	}

	db := database.GetDatabase()
	err = db.Create(&book).Error

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "could not create book",
		})

		return
	}

	ctx.JSON(201, book)
}

func ListBooks(ctx *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book
	err := db.Find(&books).Error

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "could not list books: " + err.Error(),
		})

		return
	}

	ctx.JSON(200, books)
}
