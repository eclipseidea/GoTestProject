package handler

import (
	"go_web_server/internal/repository/store"
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repos *store.Repository
}

func NewHandler(repo *store.Repository) *Handler {
	return &Handler{repos: repo}
}

func (h *Handler) InitHTTPRouter() *gin.Engine {
	router := gin.New()

	router.POST(UserAddRoute, h.AddUser)
	router.GET(UsersFindAllRoute, h.FindAllUsers)
	router.POST(UserAddBookRoute, h.AddBookToUser)
	router.DELETE(UserDeleteRoute, h.DeleteUser)
	router.DELETE(UserDeleteBookRoute, h.DeleteBookFromUser)
	router.PUT(UserUpdateRoute, h.UpdateUser)
	router.GET(UserFindByIDRoute, h.FindUserByID)

	router.POST(BookAddRoute, h.AddBook)
	router.PUT(BookUpdateRoute, h.UpdateBook)
	router.DELETE(BookDeleteRoute, h.DeleteBook)
	router.GET(BooksFindAllRoute, h.FindAllBook)
	router.GET(BookFindByNameRoute, h.FindBookByName)

	log.Println("server started on port: 8080")

	return router
}
