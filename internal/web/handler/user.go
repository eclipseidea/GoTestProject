package handler

import (
	"go_web_server/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserBooks struct {
	UserID int `json:"userId" binding:"required"`
	BookID int `json:"bookId" binding:"required"`
}

func (h *Handler) AddUser(c *gin.Context) {
	var data model.User
	if err := c.BindJSON(&data); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.repos.UserData.AddUserRepo(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) FindAllUsers(c *gin.Context) {
	list, err := h.repos.UserData.FindAllUsersRepo()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"userList": list,
	})
}

func (h *Handler) AddBookToUser(c *gin.Context) {
	var data UserBooks
	if err := c.BindJSON(&data); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.repos.UserData.UserAddBookRepo(data.UserID, data.BookID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) DeleteBookFromUser(c *gin.Context) {
	userID := c.Param("uid")
	bookID := c.Param("bid")

	uid, err := strconv.Atoi(userID)
	if err != nil || uid == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid input userID")
		return
	}

	bid, err := strconv.Atoi(bookID)
	if err != nil || uid == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid input bookID")
		return
	}

	err = h.repos.DeleteBookFromUserRepo(uid, bid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var data model.User

	if err := c.BindJSON(&data); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	affectedRow, err := h.repos.UpdateUserRepo(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"affected_row": affectedRow,
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil || id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.repos.DeleteUserRepo(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) FindUserByID(c *gin.Context) {
	userID := c.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil || id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	user, err := h.repos.FindUserByIDRepo(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"current_user": user,
	})
}
