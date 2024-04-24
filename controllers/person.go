package controllers

import (
	"latihan/database"
	"latihan/repository"
	"latihan/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllPerson(c *gin.Context) {
	var (
		result gin.H
	)

	persons, err := repository.GetAllPerson(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"results": persons,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertPerson(c *gin.Context) {
	var person structs.Person

	err := c.ShouldBindJSON(&person)
	if err != nil {
		panic(err)
	}
	person.ID = uuid.New().String()
	err = repository.InsertPerson(database.DbConnection, person)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "success insert person",
	})
}

func UpdatePerson(c *gin.Context) {
	var person structs.Person

	err := c.ShouldBindJSON(&person)
	if err != nil {
		panic(err)
	}
	person.ID = c.Param("id")

	err = repository.UpdatePerson(database.DbConnection, person)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success update person",
	})
}

func DeletePerson(c *gin.Context) {
	var person structs.Person

	person.ID = c.Param("id")

	err := repository.DeletePerson(database.DbConnection, person)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success delete person",
	})
}
