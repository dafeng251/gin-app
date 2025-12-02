package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model        // 包含了 ID, CreatedAt, UpdatedAt, DeletedAt 字段
	Username   string `gorm:"type:varchar(20);not null;comment:用户名" json:"username"`
	Password   string `gorm:"type:varchar(100);not null;comment:密码" json:"-"` // json:"-" 表示返回前端时不显示密码
	Nickname   string `gorm:"type:varchar(100);comment:昵称" json:"nickname"`
	Role       string `gorm:"type:varchar(10);default:user;comment:角色" json:"role"`
	Gender     string `gorm:"type:int(2);default:0;commit:性别" json:"gender"`
	Avatar     string `gorm:"type:varchar(200);comment:头像" json:"avatar"`
	Email      string `gorm:"type:varchar(200);commit:邮箱" json:"email"`
	Phone      string `gorm:"type:varchar(200);commit:手机号" json:"phone"`
}
