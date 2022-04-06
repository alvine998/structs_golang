package handler

import (
	"fmt"
	"net/http"
	"structs/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse
	for _, b := range books {
		bookResponse := book.BookResponse{
			ID:          b.ID,
			Title:       b.Title,
			Price:       b.Price,
			Description: b.Description,
			Rating:      b.Rating,
			Discount:    b.Discount,
		}

		booksResponse = append(booksResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "Balajar golang bareng agung setiawan",
	})
}

// URL parameter
func (h *bookHandler) BooksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// URL Double Parameter
func (h *bookHandler) BooksHandler2(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

// URL Parameter Query
func (h *bookHandler) QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

// URL Double Parameter Query
func (h *bookHandler) QueryHandler2(ctx *gin.Context) {
	title := ctx.Query("title")
	price := ctx.Query("price")
	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

// Function Post Books
func (h *bookHandler) PostBooksHandler(ctx *gin.Context) {
	var bookRequest book.BookRequest
	err := ctx.ShouldBindJSON(&bookRequest)

	// Error validation
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
