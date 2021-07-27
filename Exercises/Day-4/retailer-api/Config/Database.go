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

func GetDB() *gorm.DB {
	return DB
}


