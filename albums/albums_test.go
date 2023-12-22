package albums

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/albums", GetAlbums)
	router.GET("/albums/album/:id", GetAlbumByID)
	router.GET("/albums/:artist", GetAlbumsByArtist)
	router.POST("/albums", PostAlbums)
	router.DELETE("/albums/:id", DeleteAlbum)

	return router
}

func TestGetEmptyAlbums(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Data Not Found.")
}

func TestGetNonExistingAlbum(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/album/12314", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Data Not Found.")
}

func TestGetAlbumWithNonValidID(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/album/dsglksg434", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Opetaion Failed.")
}

func TestPostAlbum(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	newAlbum := album{
		Title:  "TestTitle",
		Artist: "TestArtist",
		Price:  500,
	}

	payload, err := json.Marshal(newAlbum)
	assert.Nil(t, err, "Error marshaling JSON")

	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(payload))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Success")
}

func TestGetAlbumByID(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/album/1", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Success")
}

func TestGetAlbums(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Success")
}

func TestGetAlbumsByArtist(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/TestArtist", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Success")
}

func TestGetAlbumsByNonExistingArtist(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/nonexistingartist", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code, "Data Not Found.")
}

func TestDeleteAlbum(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/albums/1", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Success")
}

func TestDeleteAlbumWithNonValidId(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/albums/test123", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Operation Failed.")
}

func TestDeleteAlbumWithNonExistingId(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/albums/8865635", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code, "Data Not Found.")
}
