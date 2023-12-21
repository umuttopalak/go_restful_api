package todos

import (
	"example/restfulapi/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var todos []todo

func PostTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		utils.NewErrorResponse(c, 400)
	}

	newTodo.ID = utils.GetNextID()
	todos = append(todos, newTodo)

	utils.NewSuccessResponse(c, newTodo)
}

func GetTodos(c *gin.Context) {
	if len(todos) > 0 {
		utils.NewSuccessResponse(c, todos)
	} else {
		utils.NewErrorResponse(c, 404)
	}
}

func GetTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.NewErrorResponse(c, 500)
	}

	for _, t := range todos {
		if t.ID == id {
			utils.NewSuccessResponse(c, t)
			return
		}
	}

	utils.NewErrorResponse(c, 404)
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
		utils.NewSuccessResponse(c, filteredTodos)
	} else {
		utils.NewErrorResponse(c, 404)
	}

}

func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.NewErrorResponse(c, 400)
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
		utils.NewSuccessResponse(c, nil)
	} else {
		utils.NewErrorResponse(c, 404)
	}
}
