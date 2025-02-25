package es_ser

import (
	"goblog_server/models"
)

type EsService struct {
}

func (s EsService) CommList() {

}

type Option struct {
	models.PageInfo
	Fields []string
	Tag    string
}
