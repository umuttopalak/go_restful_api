package main

import (
	"example/restfulapi/albums"
	"example/restfulapi/docs"
	"example/restfulapi/todos"
	"example/restfulapi/users"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		album := v1.Group("/albums")
		{
			album.GET("/", albums.GetAlbums)
			album.GET("/album/:id", albums.GetAlbumByID)
			album.GET("/:artist", albums.GetAlbumsByArtist)
			album.POST("/", albums.PostAlbums)
			album.DELETE("/:id", albums.DeleteAlbum)
		}
		todo := v1.Group("/todos")
		{
			todo.POST("/", todos.PostTodo)
			todo.GET("/", todos.GetTodos)
			todo.GET("/todo/:id", todos.GetTodoById)
			todo.GET("/:author", todos.GetTodoByAuthor)
			todo.DELETE("/:id", todos.DeleteTodo)
		}
		user := v1.Group("/users")
		{
			user.POST("/", users.PostUser)
			user.GET("/", users.GetUsers)
			user.GET("/user/:id", users.GetUser)
			user.DELETE("/user/:id", users.DeleteUser)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
