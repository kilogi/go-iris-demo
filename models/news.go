package models

import (
	"time"
)

//表名
func (News) TableName() string {
	return "news"
}

//表迁移
type News struct {
	ID        uint       `gorm:"primary_key" json:"id"`  //这里不使用json:"xxx"的话，输出字段会为大写
	Title     string     `gorm:"size:20" json:"title"`
	Intro     string     `grom:"size:150" json:"intro"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//验证字段规则
type NewsJson struct {
	Title string `json:"title" validate:"required,gte=2,lte=20"`
	Intro string `json:"intro" validate:"required,gte=10,lte=150"`
}
