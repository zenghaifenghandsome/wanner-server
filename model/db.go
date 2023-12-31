package model

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dbConfig := "host=" + viper.GetString("database.DbHost") +
		" user=" + viper.GetString("database.DbUser") +
		" password=" + viper.GetString("database.DbPassword") +
		" dbname=" + viper.GetString("database.DbName") +
		" port=" + viper.GetString("database.DbPort") +
		" sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dbConfig), &gorm.Config{})

	if err != nil {
		fmt.Println(dbConfig)
		fmt.Println("连接数据库失败，请检查参数：", err)
		panic(err)
	} else {
		fmt.Println("连接数据库 成功")
		resErr := db.AutoMigrate(
			&User{},
		)
		if resErr != nil {
			fmt.Println(resErr)
		}

		sqlDB, _ := db.DB()

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxIdleTime(10 * time.Hour)

	}
}
