package albums

import (
	"example/restfulapi/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var albums []album

func GetAlbums(c *gin.Context) {
	if len(albums) > 0 {
		utils.NewSuccessResponse(c, albums)
	} else {
		utils.NewErrorResponse(c, 404)
	}
}

func GetAlbumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.NewErrorResponse(c, 500)
	}

	for _, a := range albums {
		if a.ID == id {
			utils.NewSuccessResponse(c, a)
			return
		}
	}

	utils.NewErrorResponse(c, 404)
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
		utils.NewSuccessResponse(c, filtered_albums)
	} else {
		utils.NewErrorResponse(c, 404)
	}

}

func PostAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		utils.NewErrorResponse(c, 400)
	}

	newAlbum.ID = utils.GetNextID()
	albums = append(albums, newAlbum)
	utils.NewSuccessResponse(c, newAlbum)
}

func DeleteAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.NewErrorResponse(c, 400)
	}

	index := -1
	for i, a := range albums {
		if a.ID == id {
			index = i
			break
		}
	}

	if index != -1 {
		albums = append(albums[:index], albums[index+1:]...)
		utils.NewSuccessResponse(c, nil)
	} else {
		utils.NewErrorResponse(c, 404)
	}

}
