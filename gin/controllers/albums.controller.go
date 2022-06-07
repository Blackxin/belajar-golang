package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tfkhdyt/belajar-golang/gin/data"
	"github.com/tfkhdyt/belajar-golang/gin/structs"
)

// getAlbums responds with the list of all albums aj JSON
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum structs.Album
	err := c.BindJSON(&newAlbum)

	if err != nil {
		log.Fatalln(err)
		return
	}

	data.Albums = append(data.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range data.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
