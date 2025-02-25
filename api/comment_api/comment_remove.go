package comment_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog_server/global"
	"goblog_server/models"
	"goblog_server/models/res"
	"goblog_server/service/redis_ser"
	"goblog_server/utils"
	"gorm.io/gorm"
)

// 删除根评论+删除子评论

func (CommentApi) CommentRemoveView(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("评论不存在", c)
		return
	}
	// 统计评论下的子评论数 再把自己算上去
	subCommentList := FindSubCommentCount(commentModel)
	count := len(subCommentList) + 1
	redis_ser.NewCommentCount().SetCount(commentModel.ArticleID, -count)

	// 判断是否是子评论
	if commentModel.ParentCommentID != nil {
		// 子评论
		// 找父评论，减掉对应的评论数
		global.DB.Model(&models.CommentModel{}).
			Where("id = ?", *commentModel.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count - ?", count))
	}

	// 删除子评论以及当前评论
	/*
		在数据库中，通常会存在 外键约束，而在评论系统中，父评论和子评论之间可能存在外键依赖。也就是说，子评论可能会引用父评论的 ID，
		这就意味着，删除操作必须遵循一定的顺序。否则，删除父评论时可能会违反数据库的外键约束，因为子评论仍然存在且引用了父评论的 ID。
		反转的目的是：
		先删除子评论，后删除父评论。这是为了避免因删除父评论时破坏子评论和父评论之间的关系，从而导致外键约束错误。
	*/
	var deleteCommentIDList []uint
	for _, model := range subCommentList {
		deleteCommentIDList = append(deleteCommentIDList, model.ID)
	}

	// 反转，然后一个一个删
	utils.Reverse(deleteCommentIDList)
	deleteCommentIDList = append(deleteCommentIDList, commentModel.ID)
	for _, id := range deleteCommentIDList {
		global.DB.Model(models.CommentModel{}).Delete("id = ?", id)
	}

	res.OkWithMessage(fmt.Sprintf("共删除 %d 条评论", len(deleteCommentIDList)), c)
	return
}
