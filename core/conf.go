package core

import (
	"fmt"
	"goblog_server/config"
	"goblog_server/global"
	"gopkg.in/yaml.v2"
	"io/fs"
	"log"
	"os"
)

const ConfigFile = "settings.yaml"

// InitConf 读取yaml配置文件
func InitConf() {

	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)

	if err != nil { // 读取配置文件失败
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil { // 反序列化失败
		log.Fatalf("config Init Unmarshal error : %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}

func WriteConf() error {
	marshal, err := yaml.Marshal(global.Config)
	if err != nil {
		//global.Log.Error(err)
		return err
	}

	err = os.WriteFile(ConfigFile, marshal, fs.ModePerm)
	if err != nil {
		//global.Log.Error(err)
		return err
	}
	global.Log.Info("ConfigFile has been modified")
	return nil
}
