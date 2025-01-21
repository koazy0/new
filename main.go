package main

import (
	"goblog_server/core"
	_ "goblog_server/docs"
	"goblog_server/flag"
	"goblog_server/global"
	"goblog_server/routers"
)

// @goblog API文档
// @version 1.0
// @description goblog API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {

	core.InitConf()                // 读取配置文件
	global.Log = core.InitLogger() // 初始化日志
	global.DB = core.InitGorm()    // 连接mysql
	//todo 这里可以优化
	// 在解析命令行参数的时候可以分别促使话
	global.Redis = core.InitRedis()    // 连接redis
	global.ESClient = core.EsConnect() // 连接ES

	// 命令行参数绑定
	Option := flag.Parse()
	if flag.IsWebStop(Option) {
		flag.SwitchOption(Option)
		return //如果需要停止（即要迁移表结构或创建用户的话），则不运行服务器，只做对应功能
	}

	global.Router = routers.InitRouter()
	global.Log.Info("Server 运行在 " + "global.Config.System.Addr()" + "上")
	err := global.Router.Run(global.Config.System.Addr())
	if err != nil {
		global.Log.Error(err)
	}
}
