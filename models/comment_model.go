package models

// CommentModel 评论表
// 评论其实是一个树形结构，有父亲有孩子
// 带上ID的 结构，要么都加上容量限制，要么都不加上
type CommentModel struct {
	MODEL
	SubComments        []*CommentModel `gorm:"foreignkey:ParentCommentID" json:"sub_comments"`  // 子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignkey:ParentCommentID" json:"comment_model"` // 父级评论
	ParentCommentID    *uint           `json:"parent_comment_id"`                               // 父评论id
	Content            string          `gorm:"size:256" json:"content"`                         // 评论内容
	DiggCount          int             `gorm:"size:8;default:0;" json:"digg_count"`             // 点赞数
	CommentCount       int             `gorm:"size:8;default:0;" json:"comment_count"`          // 子评论数
	ArticleESID        string          `json:"article_es_id"`                                   // 文章id
	User               UserModel       `json:"user"`                                            //关联的用户
	UserID             uint            `json:"user_id"`                                         // 评论的用户
}
