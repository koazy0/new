package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"goblog_server/global"
	"goblog_server/models"
	"goblog_server/models/res"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")

	// 先找寻到上传参数id对应的菜单，并把先前的banner清空
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	// menuModel中的[]Banners指定了关联的中间表  many2many:menu_banner_models
	// Clear() 方法会删除菜单与所有banner之间的关系,但不会删除banner本身
	// 若要删除banner本身，应该调用delete()方法
	global.DB.Model(&menuModel).Association("Banners").Clear()

	// 如果选择了banner，那就添加
	if len(cr.ImageSortList) > 0 {
		// 操作第三张表
		var bannerList []models.MenuImageModel
		for _, sort := range cr.ImageSortList {
			// 传来的ImageSortList包含了多张图片的ID和Sort
			bannerList = append(bannerList, models.MenuImageModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("创建菜单图片失败", c)
			return
		}
	}

	// 普通更新
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}

	res.OkWithMessage("修改菜单成功", c)

}
