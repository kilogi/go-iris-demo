package services

import (
	"fmt"
	"go_mod/models"
	"golang.org/x/crypto/bcrypt"
)

//用户登录验证
func Login(name string, password string) (userInfo map[string]interface{}) {

	//验证
	user := models.Users{}
	UniteDB().Where("name=?", name).First(&user)
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		fmt.Println("密码错误")
		return
	}
	//返回信息
	userInfo = make(map[string]interface{})
	userInfo["id"] = user.ID
	userInfo["name"] = user.Name
	userInfo["gender"] = user.Gender

	return userInfo
}

//用户创建
func CreateUser(user *models.Users) {
	if err := UniteDB().Create(user).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
	}

	return
}
