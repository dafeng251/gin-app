package global

import (
	"gin-app/config"
	"gorm.io/gorm"
)

var (
	CONF config.Config // 全局配置对象
	DB   *gorm.DB      // 全局数据库连接对象
)
