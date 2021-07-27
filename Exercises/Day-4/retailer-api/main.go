package main

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Config"
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Routes"
	"log"
)


func main() {
	Config.SetupDatabase()
	router := Routes.SetupRouter()
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
