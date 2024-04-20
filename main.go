package main

import (
	"database/sql"
	"fmt"
	"practice/controllers"
	"practice/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load env")
	} else {
		fmt.Println("success load env")
	}
	psqlInfo := fmt.Sprintf("host=#{host} port=#{port} user=#{user} dbname=#{dbname} sslmode=disable")

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB connection failed")
		panic(err)
	} else {
		fmt.Println("DB connection success")
	}

	database.DbMigrate(DB)
	defer DB.Close()

	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	router.Run("localhost:8080")
}
