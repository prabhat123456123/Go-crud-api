package controllers

import (
	"net/http"
	"strconv"

	"go-crud-api/models"
	"go-crud-api/services"
	"go-crud-api/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.APIResponse{
			Status: false,
			Message: err.Error(),
		})
		return
	}

	if err := services.CreateUser(&user); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, utils.APIResponse{
		Status: true,
		Message: "User created successfully",
		Data: user,
	})
}


func GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	users, total, err := services.GetUsers(page, limit)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse{
		Status: true,
		Message: "Users fetched successfully",
		Data: gin.H{
			"records": users,
			"page": page,
			"limit": limit,
			"total": total,
		},
	})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := services.GetUserByID(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse{
		Status: true,
		Message: "User fetched successfully",
		Data: user,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, utils.APIResponse{
			Status: false,
			Message: err.Error(),
		})
		return
	}

	user, err := services.UpdateUser(id, &updatedUser)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse{
		Status: true,
		Message: "User updated successfully",
		Data: user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteUser(id); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, utils.APIResponse{
		Status: true,
		Message: "User deleted successfully",
	})
}
