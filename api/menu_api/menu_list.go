package menu_api

import (
	"github.com/gin-gonic/gin"
	"goblog_server/global"
	"goblog_server/models"
	"goblog_server/models/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (MenuApi) MenuListView(c *gin.Context) {
	// 先查菜单
	var menuList []models.MenuModel
	var menuIDList []uint

	// menuList存储所有menu menuModel
	// menuIDList存储所有menu的ID
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	// 查连接表
	var menuBanners []models.MenuImageModel

	//最终构造的结果如下所示
	//[
	//  {
	//    MenuModel: { ID: 1, MenuTitle: "首页", Sort: 10 },
	//    Banners: [
	//      { ID: 1, Path: "/img1" },
	//      { ID: 2, Path: "/img2" }
	//    ]
	//  },
	//  {
	//    MenuModel: { ID: 2, MenuTitle: "产品", Sort: 8 },
	//    Banners: [
	//      { ID: 3, Path: "/img3" }
	//    ]
	//  }
	//]

	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDList)
	var menus []MenuResponse
	for _, menuModel := range menuList {
		// menuModel就是一个菜单
		// 对每个Menu，找寻它所有关联的图片
		var banners []Banner
		for _, menuBannerModel := range menuBanners {
			if menuModel.ID != menuBannerModel.MenuID {
				continue
			}

			// 如果对应的MenuID相等，说明菜单和图片相关联，将该菜单关联的图片切片里添加该图片
			banners = append(banners, Banner{
				ID:   menuBannerModel.BannerID,
				Path: menuBannerModel.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: menuModel,
			Banners:   banners,
		})
	}
	res.OkWithData(menus, c)
	return
}
