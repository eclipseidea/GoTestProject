package handler

import (
	"go_web_server/internal/repository/store"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

func TestInitHTTPRouter(t *testing.T) {
	router := gin.New()

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			return
		}
	}()

	repos := store.NewRepository(db)
	h := NewHandler(repos)

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

	h.InitHTTPRouter()
}
