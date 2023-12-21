package users

import (
	"example/restfulapi/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users []user

func GetUsers(c *gin.Context) {
	if len(users) > 0 {
		utils.NewSuccessResponse(c, users)
	} else {
		utils.NewErrorResponse(c, 404)
	}
}

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

func PostUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		utils.NewErrorResponse(c, 400)
	}

	newUser.ID = utils.GetNextID()
	users = append(users, newUser)
	utils.NewSuccessResponse(c, newUser)
}

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
