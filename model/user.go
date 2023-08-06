package model

import "gorm.io/gorm"

// 用户基本信息表-注册表
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Phone    string `gorm:"type:varchar(20);not null" json:"phone"`
	Email    string `gorm:"type:varchar(30)" json:"email"`
	Avater   string `gorm:"type:varchar(200);default:''" json:"avater"`
	NickName string `gorm:"type:varchar(20)" json:"nickname"`
	Decrib   string `gorm:"type:varchar(200)" json:"decrib"`
	Addr     string `gorm:"type:varchar(50)" json:"addr"`
	QQ       string `gorm:"type:varchar(12)" json:"QQ"`
	Wechat   string `gorm:"type:varchar(40)" json:"wechat"`
	Role     string `gorm:"type:varchar(10);not null" json:"role"`
}
