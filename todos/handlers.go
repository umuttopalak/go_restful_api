package todos

import (
	"example/restfulapi/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var todos []todo

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// @Summary Creates a new todo
// @Description Creates a new todo with the provided data
// @Tags todos
// @Accept json
// @Produce json
// @Param newTodo body todo true "New todo data"
// @Success 200 {object} todo
// @Success 400 {object} ErrorResponse
// @Router /todos [post]
func PostTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		utils.NewErrorResponse(c, 400)
	}

	newTodo.ID = utils.GetNextID()
	todos = append(todos, newTodo)

	utils.NewSuccessResponse(c, newTodo)
}

// @Summary Returns todos
// @Description Returns a list of todos
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} todo "List of todos"
// @Success 404 {object} ErrorResponse
// @Router /todos [get]
func GetTodos(c *gin.Context) {
	if len(todos) > 0 {
		utils.NewSuccessResponse(c, todos)
	} else {
		utils.NewErrorResponse(c, 404)
	}
}

// @Summary Returns single todo
// @Description Returns a single todo by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID" Format(int64)
// @Success 200 {object} todo
// @Success 404 {object} ErrorResponse
// @Router /todos/todo/{id} [get]
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

// @Summary Returns todos by author
// @Description Returns a list of todos by a specific author
// @Tags todos
// @Accept json
// @Produce json
// @Param author path string true "Author name" Format(string)
// @Success 200 {array} todo "List of todos"
// @Success 404 {object} ErrorResponse
// @Router /todos/{author} [get]
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

// @Summary Deletes a todo
// @Description Deletes a todo with the specified ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID" Format(int64)
// @Success 200 {object} todo
// @Success 404 {object} ErrorResponse
// @Router /todos/{id} [delete]
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
