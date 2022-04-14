package handler

import (
	"go_web_server/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddBook(c *gin.Context) {
	var data model.Book
	if err := c.BindJSON(&data); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.repos.AddBookRepo(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"book": id,
	})
}

func (h *Handler) UpdateBook(c *gin.Context) {
	var data model.Book
	if err := c.BindJSON(&data); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	affectedRow, err := h.repos.UpdateBookRepo(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"affected_row": affectedRow,
	})
}

func (h *Handler) DeleteBook(c *gin.Context) {
	bookID := c.Param("id")

	id, err := strconv.Atoi(bookID)
	if err != nil || id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.repos.DeleteBookRepo(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) FindAllBook(c *gin.Context) {
	books, err := h.repos.BookData.FindAllBooksRepo()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"book list": books,
	})
}

func (h *Handler) FindBookByName(c *gin.Context) {
	bookName := c.Param("name")

	book, err := h.repos.FindBookByNameRepo(bookName)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"current_book": book,
	})
}
