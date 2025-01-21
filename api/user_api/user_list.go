package user_api

import (
	"github.com/gin-gonic/gin"
	"goblog_server/models"
	"goblog_server/models/ctype"
	"goblog_server/models/res"
	"goblog_server/service/common"
	"goblog_server/utils/desense"
	"goblog_server/utils/jwts"
)

func (UserApi) UserListView(c *gin.Context) {
	// todo 问gpt修改前的代码和之后的代码有何不同

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	// 如何判断是管理员
	//token := c.Request.Header.Get("token")
	//if token == "" {
	//	res.FailWithMessage("未携带token", c)
	//	return
	//}
	//claims, err := jwts.ParseToken(token)
	//if err != nil {
	//	res.FailWithMessage("token错误", c)
	//	return
	//}
	// 用中间件jwt_auth 进行处理
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 管理员
			user.UserName = ""
		}
		user.Tel = desense.DesensitizationTel(user.Tel)
		user.Email = desense.DesensitizationEmail(user.Email)
		// 脱敏
		users = append(users, user)
	}

	res.OKWithList(users, count, c)
}
