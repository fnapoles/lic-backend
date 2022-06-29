package main

import (
	"net/http"

	"github.com/fnapoles/lic-backend/entities"
	"github.com/gin-gonic/gin"
)

var files = []entities.File {
    {ID: "1", Title: "El arrepentimiento I", Description: "Descripcion de un verdadero arrepentimiento I"},
    {ID: "2", Title: "El arrepentimiento II", Description: "Descripcion de un verdadero arrepentimiento II"},
    {ID: "3", Title: "El arrepentimiento III", Description: "Descripcion de un verdadero arrepentimiento III"},
}

func main() {
    router := gin.Default()
    router.GET("/files", getFiles)
    router.GET("/files/:id", getFileByID)
    router.POST("/files", postFiles)

    router.Run("localhost:3000")
}

func getFiles(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, files)
}

func postFiles(c *gin.Context) {
    var newFile entities.File

    if err := c.BindJSON(&newFile); err != nil {
        return
    }

    files = append(files, newFile)
    c.IndentedJSON(http.StatusCreated, newFile)
}

func getFileByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range files {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "File not found"})
}