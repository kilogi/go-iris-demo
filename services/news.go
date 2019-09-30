package services

import (
	"fmt"
	"go_mod/models"
)

//获取新闻列表
func GetNewsList(page int, limits int) (list interface{}) {
	offset := (page * limits) - limits
	list = UniteDB().Limit(limits).Offset(offset).Find(&models.News{})
	return
}

//新建新闻
func CreateNews(news *models.News) {
	if err := UniteDB().Create(news).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
	}
	return
}

//新闻详情
func NewsDetail(id int) (detail models.News) {

	UniteDB().Where("id=?", id).First(&detail)

	return detail
}
