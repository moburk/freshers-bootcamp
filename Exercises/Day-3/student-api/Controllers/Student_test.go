package Controllers

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-3/student-api/Models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var getStudentsMock func(c *gin.Context)

type studentControllerMock struct {}

func (s studentControllerMock) getStudents(c *gin.Context){
	getStudentsMock(c)
}


func TestGetStudents(t *testing.T) {
	studentCaller = studentControllerMock{}

	//Test 1
	getStudentsMock = func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	GetStudents(c)
	assert.Equal(t, http.StatusNotFound, w.Code)

	//Test 2
	studentExample := []Models.Student{{
		FirstName : "hello",
		LastName: "world",
		DOB: "01-01-2010",
		Address: "India",
	}}
	getStudentsMock = func(c *gin.Context) {
		c.JSON(http.StatusOK, studentExample)
	}
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	GetStudents(c)
	assert.Equal(t, http.StatusOK, w.Code)
	//require.JSONEq(t, string(studentExample), )
	//log.Println(c)

}