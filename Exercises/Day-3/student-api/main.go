package main

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-3/student-api/Config"
	"Github-Projects/freshers-bootcamp/Exercises/Day-3/student-api/Models"
	"Github-Projects/freshers-bootcamp/Exercises/Day-3/student-api/Routes"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error
func main() {
	Config.DB, err = gorm.Open( mysql.Open(Config.DbURL(Config.BuildDBConfig())),
		&gorm.Config{ DisableForeignKeyConstraintWhenMigrating: true })
	if err != nil {
		fmt.Println("Status:", err)
	}
	//defer Config.DB.Close() // removed in gorm v2
	Config.DB.AutoMigrate(&Models.Student{}, &Models.Subject{}, &Models.Marks{})
	r := Routes.SetupRouter()
	//running
	r.Run()
	//Controllers.SeedSubjects(Config.DB)
}

