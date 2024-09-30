package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/alandavd/empomio/internal/core/services"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) CreateUser(ctx *gin.Context) {
	_, err := h.userService.CreateUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"message": "error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"error": false,
		"message": "New user created successfully",
	})
}