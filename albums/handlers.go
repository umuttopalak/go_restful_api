package albums

import (
	"example/restfulapi/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var albums []album

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// @Summary Returns Albums
// @Description Returns a list of albums
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {array} album "List of albums"
// @Success 400 {object} ErrorResponse
// @Success 404 {object} ErrorResponse
// @Router /albums/ [get]
func GetAlbums(c *gin.Context) {
	if len(albums) > 0 {
		utils.NewSuccessResponse(c, albums)
	} else {
		utils.NewErrorResponse(c, 404)
	}
}

// @Summary Returns single album
// @Description Returns a single album by ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path int true "Album ID" Format(int64)
// @Success 200 {object} album
// @Success 400 {object} ErrorResponse
// @Success 404 {object} ErrorResponse
// @Router /albums/album/{id} [get]
func GetAlbumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.NewErrorResponse(c, 400)
	}

	for _, a := range albums {
		if a.ID == id {
			utils.NewSuccessResponse(c, a)
			return
		}
	}

	utils.NewErrorResponse(c, 404)
}

// @Summary Returns albums by artist
// @Description Returns a list of albums by a specific artist
// @Tags albums
// @Accept json
// @Produce json
// @Param artist path string true "Artist name" Format(string)
// @Success 200 {array} album "List of albums"
// @Success 404 {object} ErrorResponse
// @Router /albums/{artist} [get]
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

// @Summary Creates a new album
// @Description Creates a new album with the provided data
// @Tags albums
// @Accept json
// @Produce json
// @Param newAlbum body album true "New album data"
// @Success 200 {object} album
// @Success 400 {object} ErrorResponse
// @Router /albums [post]
func PostAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		utils.NewErrorResponse(c, 400)
	}

	newAlbum.ID = utils.GetNextID()
	albums = append(albums, newAlbum)
	utils.NewSuccessResponse(c, newAlbum)
}

// @Summary Deletes an album
// @Description Deletes an album with the specified ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path int true "Album ID" Format(int64)
// @Success 200 {object} album
// @Success 404 {object} ErrorResponse
// @Router /albums/{id} [delete]
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
		return
	} else {
		utils.NewErrorResponse(c, 404)
		return
	}
}
