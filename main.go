package main

import (
	"example/restfulapi/albums"
	"example/restfulapi/todos"
	"example/restfulapi/users"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//Albums
	router.GET("/albums", albums.GetAlbums)
	router.GET("/albums/album/:id", albums.GetAlbumByID)
	router.GET("/albums/:artist", albums.GetAlbumsByArtist)
	router.POST("/albums", albums.PostAlbums)
	router.DELETE("/albums/:id", albums.DeleteAlbum)

	//Todos
	router.POST("/todos", todos.PostTodo)
	router.GET("/todos", todos.GetTodos)
	router.GET("/todos/todo/:id", todos.GetTodoById)
	router.GET("/todos/:author", todos.GetTodoByAuthor)
	router.DELETE("/todos/:id", todos.DeleteTodo)

	//Users
	router.POST("/users", users.PostUser)
	router.GET("/users", users.GetUsers)
	router.GET("/users/user/:id", users.GetUser)
	router.DELETE("/users/user/:id", users.DeleteUser)

	router.Run("localhost:8080")
}
