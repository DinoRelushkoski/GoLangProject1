package main

import (
	"fmt"
	"strconv"

	"BookWebApi/db"
	"BookWebApi/models"
	"github.com/gin-gonic/gin"
)

type JsonResponse map[string]interface{}

func main() {
	db.Init()
	router := gin.Default()

	router.POST("/books", createBook)
	router.GET("/books", fetchAllBooks)
	router.GET("/books/:id", fetchBookByID)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)

	if err := router.Run(":8088"); err != nil {
		fmt.Println("SERVER exception")
		fmt.Println(err)
	}
}

func createBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(400, JsonResponse{
			"message": "Invalid book data",
		})
		return
	}

	if err := newBook.Save(); err != nil {
		c.JSON(400, JsonResponse{
			"message": "Failed to save the book",
		})
		return
	}

	c.JSON(200, JsonResponse{
		"message": "Book created successfully",
		"data":    newBook,
	})
}

func fetchAllBooks(c *gin.Context) {
	books, err := models.GetAllBooks()
	if err != nil {
		c.JSON(400, JsonResponse{
			"message": "Unable to retrieve books",
		})
		return
	}

	c.JSON(200, JsonResponse{
		"message": "Books retrieved successfully",
		"data":    books,
	})
}

func fetchBookByID(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, JsonResponse{
			"message": "Invalid book ID format",
		})
		return
	}

	book, err := models.GetBookById(int64(bookID))
	if err != nil {
		c.JSON(404, JsonResponse{
			"message": "Book not found",
		})
		return
	}

	c.JSON(200, JsonResponse{
		"message": "Book retrieved successfully",
		"data":    book,
	})
}

func updateBook(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, JsonResponse{
			"message": "Invalid book ID format",
		})
		return
	}

	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(400, JsonResponse{
			"message": "Invalid book data",
		})
		return
	}

	updatedBook.Id = int64(bookID)
	if err := models.UpdateBook(updatedBook); err != nil {
		c.JSON(400, JsonResponse{
			"message": "Failed to update the book",
		})
		return
	}

	c.JSON(200, JsonResponse{
		"message": "Book updated successfully",
		"data":    updatedBook,
	})
}

func deleteBook(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, JsonResponse{
			"message": "Invalid book ID format",
		})
		return
	}

	if err := models.DeleteBook(int64(bookID)); err != nil {
		c.JSON(400, JsonResponse{
			"message": "Failed to delete the book",
		})
		return
	}

	c.JSON(200, JsonResponse{
		"message": "Book deleted successfully",
	})
}
