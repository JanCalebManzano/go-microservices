package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHello() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello")
	}
}
