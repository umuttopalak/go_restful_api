package albums

import (
	"example/restfulapi/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var albums []album

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Data Not Found."})
}

func GetAlbumsByArtist(c *gin.Context) {
	artist := c.Param("artist")

	filtered_albums := []album{}

	for _, a := range albums {
		if strings.EqualFold(artist, a.Artist) {
			filtered_albums = append(filtered_albums, a)
		}
	}

	if len(filtered_albums) > 0 {
		c.IndentedJSON(http.StatusOK, filtered_albums)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Data Not Found."})
	}
}

func PostAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	newAlbum.ID = utils.GetNextID()
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
