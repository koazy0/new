package user_api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"goblog_server/global"
	"goblog_server/models"
	"goblog_server/models/res"
	"goblog_server/plugnis/email"
	"goblog_server/utils/jwts"
	"goblog_server/utils/pwd"
	"goblog_server/utils/random"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	// 用户绑定邮箱， 第一次输入是 邮箱
	// 后台会给这个邮箱发验证码
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	session := sessions.Default(c)
	if cr.Code == nil {
		// 第一次，后台发验证码
		// 生成4位验证码， 将生成的验证码存入session，用来保证邮箱验证
		code := random.Code(4)
		// 写入session
		session.Set("valid_code", code)
		err = session.Save() //一定要保存一下
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("session错误", c)
			return
		}
		err = email.NewCode().Send(cr.Email, "你的验证码是 "+code)
		if err != nil {
			global.Log.Error(err)
		}
		res.OkWithMessage("验证码已发送，请查收", c)
		return
	}
	// 第二次，用户输入邮箱，验证码，密码
	code := session.Get("valid_code")
	// 校验验证码
	if code != *cr.Code {
		res.FailWithMessage("验证码错误", c)
		return
	}
	// 修改用户的邮箱
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	// 简单校验下用户密码长度
	if len(cr.Password) < 4 {
		res.FailWithMessage("密码强度太低", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Password)
	// 第一次的邮箱，和第二次的邮箱也要做一致性校验
	// todo 也可以将第一次的邮箱写入session中，从而实现校验
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("绑定邮箱失败", c)
		return
	}
	// 完成绑定
	res.OkWithMessage("邮箱绑定成功", c)
	return
}
