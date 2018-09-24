package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	DBCon *gorm.DB
)

func init() {

	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatal("Error loading .env file")
	}

	DbUser := os.Getenv("DBUSER")
	DbPassword := os.Getenv("DBPASS")
	ConnectionString := "tcp(" + os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT") + ")"
	DbName := "gin_gorm_test"

	fmt.Printf("%s:%s@%s/%s?charset=utf8&parseTime=True\n", DbUser, DbPassword, ConnectionString, DbName)
	var err error
	DBCon, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True", DbUser, DbPassword, ConnectionString, DbName))

	if err != nil {
		panic(err)
	}

	DBCon.LogMode(true)
}
