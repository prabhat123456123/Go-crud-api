package controllers

import (
	"net/http"

	"go-crud-api/services"
	"go-crud-api/utils"
	"go-crud-api/dto"

	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.APIResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	user, err := services.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.APIResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.APIResponse{
			Status:  false,
			Message: "Token generation failed",
		})
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse{
		Status:  true,
		Message: "Login successful",
		Data: gin.H{
			"token": token,
			"user":  user,
		},
	})
}
