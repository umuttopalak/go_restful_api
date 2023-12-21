package todos

import (
	"example/restfulapi/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var todos []todo

func PostTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	newTodo.ID = utils.GetNextID()
	todos = append(todos, newTodo)

	c.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	for _, t := range todos {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Data Not Found."})
}

func GetTodoByAuthor(c *gin.Context) {
	author := c.Param("author")

	filteredTodos := []todo{}

	for _, t := range todos {
		if strings.EqualFold(t.Author, author) {
			filteredTodos = append(filteredTodos, t)
		}
	}

	if len(filteredTodos) > 0 {
		c.IndentedJSON(http.StatusOK, filteredTodos)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Data Not Found."})
	}

}

func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	index := -1

	for i, t := range todos {
		if t.ID == id {
			index = i
			break
		}
	}

	if index != -1 {
		todos = append(todos[:index], todos[index+1:]...)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}
}
