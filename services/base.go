package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//统一DB
func UniteDB() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:root@/go_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	db.LogMode(true)
	//defer db.Close()
	return db
}
