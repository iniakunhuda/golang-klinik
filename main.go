package main

import (
	"fmt"
	"golang-klinik/Config"
	"golang-klinik/Models"
	"golang-klinik/Routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildConfig()))

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Sukses!")
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetRoutes()
	r.Run("127.0.0.1:8081")

}
