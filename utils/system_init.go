package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func InitMySQL() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, //慢SQL的阈值
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	openLink := viper.GetString("mysql.name") + ":" + viper.GetString("mysql.passwd") + "@" +
		viper.GetString("mysql.link") + "/" + viper.GetString("mysql.dbname") + "?" + viper.GetString("mysql.others")
	fmt.Println(openLink)
	db, err := gorm.Open(mysql.Open(openLink), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	fmt.Println("数据库连接成功！")
}
