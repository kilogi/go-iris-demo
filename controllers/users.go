package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"go_mod/models"
	"go_mod/services"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

//登录
func Login(ctx iris.Context) {
	name := ctx.PostValue("name")
	password := ctx.PostValue("password")

	userInfo := services.Login(name, password)

	ctx.JSON(ApiResponse(SUCCESS, "登录成功", userInfo))
}

//注册
func UserRegister(ctx iris.Context) {
	name := ctx.PostValue("name")
	password := ctx.PostValue("password")
	gender, _ := ctx.PostValueInt("gender")
	//验证
	validateUser := models.UsersJson{Name: name, Password: password, Gender: gender}
	validate = validator.New()
	err := validate.Struct(validateUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	//保存
	hashPassword, _:= bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.Users{Name: name, Password: hashPassword, Gender: gender}
	userInfo:=services.CreateUser(&user)

	ctx.JSON(ApiResponse(SUCCESS, "用户创建成功", userInfo))
}

