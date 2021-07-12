package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/praveeenrm/todoapp/config"
	"github.com/praveeenrm/todoapp/models"
	"github.com/praveeenrm/todoapp/routes"
)

func main() {
	fmt.Println("Starting the server")

	config.DB, config.Err = gorm.Open("mysql", config.DBUrl(config.BuildDBConfig()))

	if config.Err != nil {
		panic(config.Err.Error())
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Todo{})

	// Set up routings
	r := routes.SetUpRoutes()

	r.Run()

}
