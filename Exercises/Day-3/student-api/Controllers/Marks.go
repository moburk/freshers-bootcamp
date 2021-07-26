package Controllers

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-3/student-api/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetMarks get all users
func GetMarks(c *gin.Context) {
	var user []Models.Marks
	err := Models.GetAllMarks(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser create user
func AddMarks(c *gin.Context) {
	var user Models.Marks
	c.BindJSON(&user)
	err := Models.CreateMarks(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetMarksByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Marks
	err := Models.GetMarksByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetMarksByStudentID Get the user by id
func GetMarksByStudentID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Marks
	err := Models.GetMarksByStudentID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetMarksBySubjectID Get the user by id
func GetMarksBySubjectID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Marks
	err := Models.GetMarksBySubjectID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateMarks update the user information
func UpdateMarks(c *gin.Context) {
	var user Models.Marks
	id := c.Params.ByName("id")
	err := Models.GetMarksByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateMarks(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//delete the user
func DeleteMarks(c *gin.Context) {
	var user Models.Marks
	id := c.Params.ByName("id")
	err := Models.DeleteMarks(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
