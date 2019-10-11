package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

//统一DB
func UniteDB() (db *gorm.DB) {
	//读取配置文件
	config:=Config("database")
	host := config.Get("Mysql.Host")
	fmt.Println(host)
	//连接数据库
	db, err := gorm.Open("mysql", "root:root@/go_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	db.LogMode(true)
	//defer db.Close()
	return db
}

//配置文件加载
func Config(configName string) (config *viper.Viper) {
	//configName 只能设置一个，否则后一个将覆盖前一个（在同一个viper实例中）
	config = viper.New()
	config.SetConfigName(configName)
	config.AddConfigPath("./config")
	config.AddConfigPath("%GOPATH/src/")
	config.SetConfigType("yaml")
	if err := config.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}

	return
}
