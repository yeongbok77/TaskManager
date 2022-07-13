package main

import (
	"github.com/yeongbok77/TaskManager/dao/mysql"
	"github.com/yeongbok77/TaskManager/dao/redis"
	"github.com/yeongbok77/TaskManager/logger"
	"github.com/yeongbok77/TaskManager/router"
	"github.com/yeongbok77/TaskManager/settings"
)

func main() {
	// 配置文件初始化
	if err := settings.Init(); err != nil {
		panic("settings.Init err")
		return
	}

	if err := logger.Init(settings.Conf.LogConfig, "dev"); err != nil {
		panic("logger.Init err")
		return
	}

	// MySQL初始化
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		panic("mysql.Init err")
		return
	}
	// Redis初始化
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		panic("redis.Init err")
		return
	}

	// 路由设置
	r := router.SetUpRouter()

	// 启动服务
	r.Run(":8080")
}
