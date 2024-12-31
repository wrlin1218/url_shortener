package controller

import (
	"github.com/wrlin1218/url_shortener/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

// 创建用户
func (uc *UserController) CreateUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if uc.UserService.CheckUserExists(c, username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	err := uc.UserService.CreateUser(c, username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// 获取用户的所有短链接
func (uc *UserController) GetAllLinksByUserName(c *gin.Context) {
	username := c.Query("username")

	links, err := uc.UserService.GetAllLinksByUserName(c, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"links": links})
}
