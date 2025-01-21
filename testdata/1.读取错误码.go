package main

import (
	"goblog_server/models/ctype"
)

const file = "models/res/errorcode.json"

func main() {
	s := ctype.PermissionUser
	json, err := s.MarshalJson()
	if err != nil {
		return
	}
	println(json)
	println(s.MarshalJson())

}
