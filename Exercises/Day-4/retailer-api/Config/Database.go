package Config

import (
	"Github-Projects/freshers-bootcamp/Exercises/Day-4/retailer-api/Models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// DBConfig represents db configuration
type dbConfigure struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func SetupDatabase() {
	var url = dbURL(buildDBConfig())
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&Models.Product{}, &Models.Customer{}, &Models.Order{})
	populateDB(db)
	DB = db
}

func buildDBConfig() *dbConfigure {
	dbConfig := dbConfigure{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "1940mysql",
		DBName:   "retailer_api",
	}
	return &dbConfig
}
func dbURL(dbConfig *dbConfigure) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func populateDB(db *gorm.DB) {
	var customers = []Models.Customer{
		{ID: "CST12345", CustomerName: "Mohammed"},
		{ID: "CST12346", CustomerName: "Gopher"},
		{ID: "CST12347", CustomerName: "Razorpay"},
		{ID: "CST12348", CustomerName: "Hello"},
		{ID: "CST12349", CustomerName: "World"},
	}
	db.Create(&customers)
	var products = []Models.Product{
		{ID: "PROD12340", ProductName: "Hammer", Price: 100, Quantity: 50},
		{ID: "PROD12341", ProductName: "Tent", Price: 1000, Quantity: 5},
		{ID: "PROD12342", ProductName: "Screwdriver", Price: 100, Quantity: 50},
		{ID: "PROD12343", ProductName: "Scooter", Price: 50000, Quantity: 10},
		{ID: "PROD12344", ProductName: "Guitar", Price: 4000, Quantity: 15},
	}
	db.Create(&products)
}

func GetDB() *gorm.DB {
	return DB
}


