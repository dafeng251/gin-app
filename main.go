package main

import (
	"fmt"
	"gin-app/global"
	"gin-app/initialize"
	"gin-app/internal/model"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Viper
	initialize.InitConfig()

	// 使用配置
	// 比如设置 Gin 的模式
	gin.SetMode(global.CONF.Server.Mode)

	// 初始化数据库 (顺序很重要，必须在配置加载之后)
	initialize.Gorm()

	// 数据库迁移 (自动建表)
	// 这一步会检查 model.User 结构体，在数据库创建 user 表
	if global.DB != nil {
		// 迁移所有需要的表
		global.DB.AutoMigrate(&model.User{})
		fmt.Println("数据库表结构迁移成功")
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// 在业务代码中直接使用全局变量获取数据库配置等
		c.JSON(200, gin.H{
			"code":    200,
			"message": "server is running",
			"data":    "hello world",
		})
	})

	// 使用配置中的端口启动
	addr := fmt.Sprintf(":%d", global.CONF.Server.Port)
	r.Run(addr)
}
