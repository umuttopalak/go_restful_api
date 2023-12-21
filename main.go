package main

import (
	"example/restfulapi/albums"
	"example/restfulapi/todos"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", albums.GetAlbums)
	router.GET("/albums/:id", albums.GetAlbumByID)
	router.POST("/albums", albums.PostAlbums)

	router.POST("/todos", todos.PostTodo)
	router.GET("/todos", todos.GetTodos)
	router.GET("/todos/todo/:id", todos.GetTodoById)
	router.GET("/todos/:author", todos.GetTodoByAuthor)

	router.Run("localhost:8080")
}
