package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go_mod/models"
	"golang.org/x/crypto/bcrypt"
	"time"
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
	//生成token
	tokenString := createToken(user.ID, user.Name)
	//返回信息
	userInfo = make(map[string]interface{})
	userInfo["name"] = user.Name
	userInfo["gender"] = user.Gender
	userInfo["token"] = tokenString

	return
}

//用户创建
func CreateUser(user *models.Users) (userInfo map[string]interface{}) {

	if err := UniteDB().Create(user).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
	}
	//生成token
	tokenString := createToken(user.ID, user.Name)
	//返回信息
	userInfo = make(map[string]interface{})
	userInfo["name"] = user.Name
	userInfo["gender"] = user.Gender
	userInfo["token"] = tokenString

	return
}

//生成token
func createToken(id uint, name string) (tokenString string) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		fmt.Printf("CreateTokenErr:%s", err)

		return
	}

	return
}
