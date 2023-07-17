package main

import (
	"gin_scaffold/bootstrap"
	"gin_scaffold/global"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化验证器
	bootstrap.InitializeValidator()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()

	// 初始化redis
	global.App.Redis = bootstrap.InitializeRedis()

	// 初始化文件系统
	bootstrap.InitializeStorage()

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 启动服务器
	bootstrap.RunServer()
}
