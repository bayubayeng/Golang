package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"web-API/book"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	var booksResponse []book.BookResponse

	books, err := h.bookService.FindAll()

	for _, value := range books {

		bookResponse := getBookResponse(value)

		booksResponse = append(booksResponse, bookResponse)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBookById(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)

	oneBook, err := h.bookService.FindByID(i)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"errors": "book is not found",
		})
	}

	convert := getBookResponse(oneBook)

	c.JSON(http.StatusOK, gin.H{

		"data": convert,
	})
}

func (h *bookHandler) PostBookHandler(c *gin.Context) {

	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, value := range err.(validator.ValidationErrors) {

			errorMessage := fmt.Sprintf("error on field %s ,condition : %s", value.Field(), value.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	h.bookService.Create(bookRequest)

	c.JSON(http.StatusOK, gin.H{

		"Data": bookRequest,
	})
}

func getBookResponse(b book.Book) book.BookResponse {

	bookResponse := book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
	}

	return bookResponse
}
