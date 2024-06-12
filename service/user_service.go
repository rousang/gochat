package service

import (
	"errors"
	"gochat/models"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
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

	if ok, err := govalidator.ValidateStruct(user); !ok {
		c.JSON(500, gin.H{
			"message": "新建用户失败",
			"Error":   err.Error(),
		})
		return
	}

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
	name := c.PostForm("name")
	user, err := models.FindUserByName(name)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Find user error",
		})
		return
	}
	// 修改内容
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	if ok, err := govalidator.ValidateStruct(user); !ok {
		c.JSON(500, gin.H{
			"message": "修改失败",
			"Error":   err.Error(),
		})
		return

	}

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

func CheckDuplicate(c *gin.Context) (bool, error) {
	name := c.Query("name")
	phone := c.Query("phone")
	email := c.Query("email")

	if _, err := models.FindUserByName(name); err == nil {
		return false, errors.New("用户名已存在")
	}

	if _, err := models.FindUserByPhone(phone); err == nil {
		return false, errors.New("手机号已存在")
	}

	if _, err := models.FindUserByEmial(email); err == nil {
		return false, errors.New("邮箱已存在")
	}

	return true, nil
}
