package flag

import "goblog_server/models"

func EsCreateIndex() {
	models.ArticleModel{}.CreateIndex()
	models.FullTextModel{}.CreateIndex()
}
