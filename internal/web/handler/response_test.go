package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_newErrorResponse(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	newErrorResponse(c, http.StatusInternalServerError, "bad request")
}
