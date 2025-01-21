package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"goblog_server/global"
	"goblog_server/utils"
)

// RouterGroup 存放group以便于路由分组
type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	// 设置gin的调试环境等级，在配置文件中更改
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	PublicGroup := router.Group("")
	PublicGroup.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.SettingsRouter() // 在组中添加系统配置API
	routerGroupApp.ImagesRouter()   // 在组中添加图片配置API
	routerGroupApp.AdvertRouter()   // 在组中添加广告配置API
	routerGroupApp.MenusRouter()    // 在族中添加菜单配置API
	routerGroupApp.UserRouter()     // 在族中添加用户配置API
	routerGroupApp.ArticleRouter()  // 在族中添加文章配置API

	// 设置静态目录
	router.GET("/static/uploads/*filepath", HandleStaticPictures)

	return router
}

func HandleStaticPictures(c *gin.Context) {
	filepath := c.Param("filepath")

	if !utils.CheckWhiteImageList(filepath, global.WhiteImageList) {
		return
	}

	// 构造实际文件路径
	fullPath := "./uploads" + filepath
	// 返回文件
	c.File(fullPath)
}
