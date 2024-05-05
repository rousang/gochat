package models

import (
	"fmt"
	"gochat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIP      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogoutTime    time.Time
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableNmae() string {
	return "user_basic"
}

// func init() {
// 	if !utils.DB.Migrator().HasTable(&UserBasic{}) {
// 		utils.DB.Migrator().CreateTable(&UserBasic{})
// 	}

// }

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user UserBasic) error {
	if err := utils.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user UserBasic) error {
	if err := utils.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user UserBasic) error {
	if err := utils.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByName(name string) (*UserBasic, error) {
	user := &UserBasic{}
	if err := utils.DB.Where("name = ?", name).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
