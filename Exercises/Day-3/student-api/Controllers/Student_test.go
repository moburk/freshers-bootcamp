package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)



func TestGetStudents(t *testing.T) {
	//r := gin.Default()
	w := httptest.NewRecorder()
	c, _ := gin.Context
	//r.GET("/student-api/student", GetStudents(c), nil)
	//r.ServeHTTP(w,req)
	GetStudents(c)
	assert.Equal(t, http.StatusOK, w.Code)
	//assert.Contains(t, "first_name", w.Body)
}