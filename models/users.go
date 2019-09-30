package models

import "time"

//表名
func (Users) TableName() string {
	return "users"
}

//表迁移
type Users struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"name:20"`
	Gender    int   `gorm:"gender:1"`
	Password  []byte `grom:"password:150"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

//验证字段规则
type UsersJson struct {
	Name     string `json:"name" validate:"required,gte=2,lte=10"`
	Gender   int   `json:"gender" validate:"required"`
	Password string `json:"password" validate:"required,gte=6,lte=16"`
}
