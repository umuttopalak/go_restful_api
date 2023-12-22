package users

import (
	"example/restfulapi/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users []user

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// @Summary Returns users
// @Description Returns a list of users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} user "List of users"
// @Success 404 {object} ErrorResponse
// @Router /users/ [get]
func GetUsers(c *gin.Context) {
	if len(users) > 0 {
		utils.NewSuccessResponse(c, users)
	} else {
		utils.NewErrorResponse(c, 404)
	}
}

// @Summary Returns single user
// @Description Returns a single user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID" Format(int64)
// @Success 200 {object} user
// @Success 404 {object} ErrorResponse
// @Router /users/user/{id} [get]
func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.NewErrorResponse(c, 400)
	}

	for _, u := range users {
		if u.ID == id {
			utils.NewSuccessResponse(c, u)
			return
		}
	}
	utils.NewErrorResponse(c, 404)
}

// @Summary Creates a new user
// @Description Creates a new user with the provided data
// @Tags users
// @Accept json
// @Produce json
// @Param newUser body user true "New user data"
// @Success 200 {object} user
// @Success 400 {object} ErrorResponse
// @Router /users/ [post]
func PostUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		utils.NewErrorResponse(c, 400)
	}

	newUser.ID = utils.GetNextID()
	users = append(users, newUser)
	utils.NewSuccessResponse(c, newUser)
}

// @Summary Deletes a user
// @Description Deletes a user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID" Format(int64)
// @Success 200 {object} user
// @Success 404 {object} ErrorResponse
// @Router /users/user/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.NewErrorResponse(c, 400)
	}

	index := -1
	for i, u := range users {
		if u.ID == id {
			index = i
			break
		}
	}

	if index != -1 {
		users = append(users[:index], users[index+1:]...)
		utils.NewSuccessResponse(c, nil)
	} else {
		utils.NewErrorResponse(c, 404)
	}
}
