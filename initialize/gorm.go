package initialize

import (
	"fmt"
	"gin-app/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Gorm 初始化数据库并产生全局变量
func Gorm() {
	m := global.CONF.Mysql // 获取之前用 Viper 读取的配置

	// 1. 拼接 DSN (Data Source Name)
	// 格式: user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		m.Username, m.Password, m.Host, m.Port, m.Dbname, m.Config,
	)

	// 2. 配置 GORM 的日志模式
	// 建议：开发环境用 Info (打印所有SQL)，生产环境用 Error (只打印错误)
	var gormLogger logger.Interface
	if global.CONF.Server.Mode == "debug" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default.LogMode(logger.Error)
	}

	// 3. 打开连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger, // 设置日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名是否单数 (User struct -> user table, 否则是 users)
		},
	})

	if err != nil {
		panic(fmt.Sprintf("数据库连接失败: %s", err))
	}

	// 4. 配置连接池 (重要！不配置会导致高并发下连接泄露或性能问题)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           // 空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          // 打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最大时间

	// 5. 赋值给全局变量
	global.DB = db

	fmt.Println("MySQL 连接成功")
}
