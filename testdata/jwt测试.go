package main

import (
	"fmt"
	"goblog_server/core"
	"goblog_server/utils/jwts"
)

func main() {

	core.InitConf() // 读取配置文件
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		Username: "lpx",
		NickName: "lpx",
		Role:     1,
		UserID:   1,
	})
	if err != nil {
		return
	}
	fmt.Println(token)
	parseToken, err := jwts.ParseToken(token)
	if err != nil {
		return
	}
	fmt.Println(parseToken)
}
