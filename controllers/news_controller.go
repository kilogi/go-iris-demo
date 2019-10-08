package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"go_mod/models"
	"go_mod/services"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func News(ctx iris.Context) {

	ctx.JSON("mvc成功")
}

//新闻列表
func NewsList(ctx iris.Context) {
	page, _ := ctx.URLParamInt("page")
	limits, _ := ctx.URLParamInt("limits")
	//获取 token 结构体
	token := ctx.Values().Get("jwt").(*jwt.Token)
	//获取jwt中信息
	fmt.Println(token.Claims.(jwt.MapClaims)["name"])
	//列表
	newsList := services.GetNewsList(page, limits)
	ctx.JSON(ApiResponse(SUCCESS, "获取新闻列表成功", newsList))
}

//新闻详情
func NewsDetail(ctx iris.Context) {
	id, _ := ctx.URLParamInt("id")
	newsDetail := services.NewsDetail(id)
	ctx.JSON(ApiResponse(SUCCESS, "详情成功", newsDetail))
}

//新建新闻
func CreateNews(ctx iris.Context) {
	title := ctx.PostValue("title")
	intro := ctx.PostValue("intro")

	validateNews := models.NewsJson{Title: title, Intro: intro}
	validate = validator.New()
	err := validate.Struct(validateNews)
	if err != nil {
		fmt.Println(err)
		return
	}

	news := models.News{Title: title, Intro: intro}
	services.CreateNews(&news)

	ctx.JSON(ApiResponse(SUCCESS, "新闻创建成功", ""))
}
