package service

import (
	"gochat/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags         用户模块
// @Summary 	 用户列表
// @Success      200  {string}  json{"code", "message"}
// @Router       /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"data": data,
	})
}

// CreateUser
// @Tags         用户模块
// @Summary 	 新增用户
// @Param        name query string true "用户名"
// @Param        name query string true "密码"
// @Success      200  {string}  json{"code", "message"}
// @Router       /user/CreateUser [get]
func CreateUser(c *gin.Context) {
	var user models.UserBasic
	// c.BindJSON(&user)
	user.Name = c.Query("name")
	user.PassWord = c.Query("password")
	user.LoginTime = time.Now()
	user.HeartbeatTime = time.Now()
	user.LogoutTime = time.Now()
	err := models.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Insert success",
	})
}

func DeleteUser(c *gin.Context) {
	var user models.UserBasic
	// user.Name = c.Query("name")
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error ID",
		})
		return
	}
	user.ID = uint(id)
	err = models.DeleteUser(user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Delete success",
	})
}

func UpdateUser(c *gin.Context) {
	// user := models.UserBasic{}
	name := c.Query("name")
	passwd := c.Query("password")
	user, err := models.GetUserByName(name)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Find user error",
		})
		return
	}
	user.PassWord = passwd
	err = models.UpdateUser(*user)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Update success",
	})
}
