package flag

import (
	sys_flag "flag"
	"github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User string // -u admin 创建管理员 -u user 创建用户
	ES   string //-es create -es delete
}

// Parse 解析命令行参数
func Parse() Option {

	// db默认值为false
	// 在命令行中传入 -db 参数，则 db 的值会变为 true
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	es := sys_flag.String("es", "", "es操作")
	// 解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
		ES:   *es,
	}
}

// IsWebStop1 是否停止web项目,已被下面的函数优化
func IsWebStop1(option Option) bool {
	// 实际上运行项目都需要关闭web项目
	// todo 先留着

	if option.DB || option.User != "" {
		return true
	}
	return false
}

// IsWebStop 是否停止web项目
// 其中有一个参数不为空则返回true
// 否则返回false
func IsWebStop(option Option) (f bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}
	}
	return f
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	if option.ES == "create" {
		EsCreateIndex()
	}
	if option.ES == "remove" {

	}
	// 如果两个条件都不成立，则说明输入有错
	//sys_flag.Usage()
}
