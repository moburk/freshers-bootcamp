package Controllers

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-3/student-api/Models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var getStudentsMock func(c *gin.Context, students *[]Models.Student) error

type studentControllerMock struct {}

func (s studentControllerMock) getStudents(c *gin.Context, students *[]Models.Student) error{
	return getStudentsMock(c, students)
}


func TestGetStudents(t *testing.T) {
	studentCaller = studentControllerMock{}

	//Test 1
	getStudentsMock = func(c *gin.Context, students *[]Models.Student) error{
		return errors.New("error encountered")
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
	getStudentsMock = func(c *gin.Context, students *[]Models.Student) error{
		students = &studentExample
		return nil
	}
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	GetStudents(c)
	assert.Equal(t, http.StatusOK, w.Code)
	//require.JSONEq(t, string(studentExample), )
	fmt.Println(w)

}