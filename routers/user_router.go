package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"goblog_server/api"
	"goblog_server/middlewares"
)

var store = cookie.NewStore([]byte("HyhSDHB8D9H2D92haS1"))

func (r RouterGroup) UserRouter() {
	user_api := api.Apigroup.UserApi
	//r.Use(sessions.Sessions("sessionID", store)) 是为当前路由组 UserRouter 添加会话中间件的配置。
	//它只会影响与这个路由组相关的请求，因此它不会直接影响到 SettingsRouter 中的路由。
	/*
		会话中间件 (sessions.Sessions) 与 JWT 认证： sessions.Sessions("sessionID", store) 主要用于会话管理，而 middlewares.JwtAuth() 主要用于基于 JWT 的身份验证。
		两者功能是不同的，且不冲突。会话中间件不会影响 JWT 中间件，反之亦然。
		注销操作中的 JWT 认证： 在 POST /logout 路由中，middlewares.JwtAuth() 只会验证 JWT 令牌的有效性。如果验证通过，用户会执行 LogoutView，此时可以根据需要清除会话、退出等操作。
	*/
	r.Use(sessions.Sessions("sessionID", store))
	r.POST("email_login", user_api.EmailLoginView)
	r.POST("users", user_api.UserCreateView) //感觉去掉jwtauth比较合适，毕竟创建用户不需要认证
	r.GET("users", middlewares.JwtAuth(), user_api.UserListView)
	r.PUT("user_role", middlewares.JwtAuth(), user_api.UserUpdateRoleView)
	r.PUT("user_password", middlewares.JwtAuth(), user_api.UserUpdatePassword)
	r.POST("logout", middlewares.JwtAuth(), user_api.LogoutView)
	r.DELETE("users", middlewares.JwtAuth(), user_api.UserRemoveView)
	r.POST("user_bind_email", middlewares.JwtAuth(), user_api.UserBindEmailView)
}
