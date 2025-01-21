package service

import (
	"goblog_server/service/image_ser"
	"goblog_server/service/user_ser"
)

type ServiceGroup struct {
	ImageService image_ser.ImageService
	UserService  user_ser.UserService
}

var ServiceApp ServiceGroup
