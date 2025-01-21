package flag

import (
	"fmt"
	"goblog_server/global"
	"goblog_server/models"
	"goblog_server/models/ctype"
	"goblog_server/utils/pwd"
)

func CreateUser(permission string) {

	// 创建用户的逻辑
	// 用户名 昵称 密码 确认密码 邮箱
	var (
		username   string
		nickname   string
		password   string
		repassword string
		email      string
	)
	fmt.Printf("请输入用户名:")
	fmt.Scan(&username)
	fmt.Printf("请输入昵称:")
	fmt.Scan(&nickname)
	// todo 合法密码校验
	// todo 位数校验GenerateFromPassword传入的参数不能超过72字节
	// todo 正则校验密码强度
	// todo 弱密码本校验
	fmt.Printf("请输入密码:")
	fmt.Scan(&password)
	fmt.Printf("请确认密码:")
	fmt.Scan(&repassword)
	fmt.Printf("请输入邮箱:")
	fmt.Scan(&email) //fmt.Scanln在除第一个输入以外的地方，后面的都会马上赋值

	// 判断用户名是否存在
	userModel := models.UserModel{}
	err := global.DB.Take(&userModel, "user_name = ?", username).Error
	if err == nil {
		// 如果存在
		global.Log.Error("用户名已存在，请重新输入")
		return
	}

	// 校验两次密码是否相等
	if password != repassword {
		global.Log.Error("两次密码不匹配，请重新输入")
		return
	}
	// 将密码存储为哈希
	hashPwd := pwd.HashPwd(password)

	// 判定权限
	role := ctype.PermissionUser
	if permission == "admin" {
		role = ctype.PermissionAdmin
	}

	// 头像问题，这里采用第一个方法
	// 1.默认头像
	// 2.随机头像
	default_avatar := "/static/uploads/avatar/default_avatar.jpg"
	//存入数据库中
	err = global.DB.Create(&models.UserModel{
		NickName:   nickname,
		UserName:   username,
		Password:   hashPwd,
		Avatar:     default_avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		Email:      email,
		Role:       role,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	global.Log.Infof("用户%s创建成功", username)
	return
}
