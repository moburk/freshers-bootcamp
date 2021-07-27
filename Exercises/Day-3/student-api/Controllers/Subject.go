package Controllers

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-3/student-api/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func SeedSubjects(db *gorm.DB) {
	subjects := []Models.Subject{
		{ SubjectName: "Math" },
		{ SubjectName: "English" },
		{ SubjectName: "Science" },
		{ SubjectName: "Conputer Science" },
	}
	db.Create(&subjects)
	//for _,subject := range subjects {
	//	subject.Id
	//
}

// Getsubjects get all subjects
func GetSubjects(c *gin.Context) {
	var subject []Models.Subject
	err := Models.GetAllSubjects(&subject)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, subject)
	}
}

// Createsubject create subject
func CreateSubject(c *gin.Context) {
	var subject Models.Subject
	c.BindJSON(&subject)
	err := Models.CreateSubject(&subject)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, subject)
	}
}
//GetsubjectByID Get the subject by id
func GetSubjectByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var subject Models.Subject
	err := Models.GetSubjectByID(&subject, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, subject)
	}
}

// Updatesubject update the subject information
func UpdateSubject(c *gin.Context) {
	var subject Models.Subject
	id := c.Params.ByName("id")
	err := Models.GetSubjectByID(&subject, id)
	if err != nil {
		c.JSON(http.StatusNotFound, subject)
	}
	c.BindJSON(&subject)
	err = Models.UpdateSubject(&subject, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, subject)
	}
}
//delete the subject
func DeleteSubject(c *gin.Context) {
	var subject Models.Subject
	id := c.Params.ByName("id")
	err := Models.DeleteSubject(&subject, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
