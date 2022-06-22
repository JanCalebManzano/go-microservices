package handlers

import (
	"net/http"
	"strconv"

	"github.com/JanCalebManzano/go-microservices/services/user/http/responses"
	"github.com/JanCalebManzano/go-microservices/services/user/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (*UserHandler) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users := models.GetUsers()
		c.JSON(http.StatusOK, users)
	}
}

func (*UserHandler) AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		in, _ := c.Get(userKey)
		u := in.(models.User)

		models.AddUser(&u)
		c.JSON(http.StatusOK, "User added")
	}
}

func (*UserHandler) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := c.Param("id")
		id, err := strconv.Atoi(sid)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
			return
		}

		in, _ := c.Get(userKey)
		u := in.(models.User)

		if err := models.UpdateUser(id, &u); err != nil {
			c.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, "User updated")
	}
}

var userKey string = "user"

func (*UserHandler) ValidateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var u models.User
		if err := c.BindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
			return
		}

		if err := u.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
			return
		}

		c.Set(userKey, u)
		c.Next()
	}
}
