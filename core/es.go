package core

import (
	"github.com/olivere/elastic/v7"
	"goblog_server/global"
)

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := global.Config.ES.URL()

	// 这里可以进行更多字段的控制
	c, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth(global.Config.ES.User, global.Config.ES.Password),
	)
	if err != nil {
		global.Log.Fatalf("es连接失败 %s", err.Error())
	}
	return c
}
