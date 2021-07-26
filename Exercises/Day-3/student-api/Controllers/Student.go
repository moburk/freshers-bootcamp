package Controllers

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-3/student-api/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type studentInterface interface {
	getStudents(c *gin.Context)
}

type StudentController struct {
}

func (student StudentController) getStudents(c *gin.Context) {
	var user []Models.Student
	err := Models.GetAllStudents(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

var studentCaller studentInterface

func init() {
	studentCaller = StudentController{}
}

//GetStudents returns all students from DB
func GetStudents(c *gin.Context) {
	studentCaller.getStudents(c)
}


//CreateStudent create student
func CreateStudent(c *gin.Context) {
	var user Models.Student
	c.BindJSON(&user)
	err := Models.CreateStudent(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//GetUserByID Get the user by id
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Student
	err := Models.GetStudentByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser update the user information
func UpdateStudent(c *gin.Context) {
	var user Models.Student
	id := c.Params.ByName("id")
	err := Models.GetStudentByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateStudent(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//delete the user
func DeleteStudent(c *gin.Context) {
	var user Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
