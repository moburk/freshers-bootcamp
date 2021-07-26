package Routes

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-3/student-api/Controllers"
	"github.com/gin-gonic/gin"
)

// configuring routes

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/student-api/students")
	{
		grp1.GET("", Controllers.GetStudents)
		grp1.POST("", Controllers.CreateStudent)
		grp1.GET("/:id", Controllers.GetStudentByID)
		grp1.PUT("/:id", Controllers.UpdateStudent)
		grp1.DELETE("/:id", Controllers.DeleteStudent)
	}
	grp2 := r.Group("/student-api/subjects")
	{
		grp2.GET("", Controllers.GetSubjects)
		grp2.POST("", Controllers.CreateSubject)
		grp2.GET("/:id", Controllers.GetSubjectByID)
		//grp1.PUT("/:id", Controllers.UpdateStudent)
		grp2.DELETE("/:id", Controllers.DeleteSubject)
	}
	grp3 := r.Group("/student-api/marks")
	{
		grp3.GET("", Controllers.GetMarks)
		grp3.POST("", Controllers.AddMarks)
		grp3.GET("/student/:id", Controllers.GetMarksByStudentID)
		grp3.GET("/subject/:id", Controllers.GetMarksBySubjectID)
		grp3.PUT("/:id", Controllers.UpdateMarks)
		grp3.DELETE("/:id", Controllers.DeleteMarks)
	}
	return r
}