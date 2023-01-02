package config

import (
	"log"

	"github.com/tigerinus/xin/common"
	"github.com/tigerinus/xin/model"
	"gopkg.in/ini.v1"
)

var (
	CommonInfo = &model.CommonModel{
		RuntimePath: "/var/run/casaos",
	}

	AppInfo = &model.APPModel{
		LogPath:     "/var/log/casaos",
		LogSaveName: common.MessageBusServiceName,
		LogFileExt:  "log",
	}

	Cfg            *ini.File
	ConfigFilePath string
)

func InitSetup(config string) {
	ConfigFilePath = MessageBusConfigFilePath
	if len(config) > 0 {
		ConfigFilePath = config
	}

	var err error

	Cfg, err = ini.Load(ConfigFilePath)
	if err != nil {
		panic(err)
	}

	mapTo("common", CommonInfo)
	mapTo("app", AppInfo)
}

func SaveSetup(config string) {
	reflectFrom("common", CommonInfo)
	reflectFrom("app", AppInfo)

	configFilePath := MessageBusConfigFilePath
	if len(config) > 0 {
		configFilePath = config
	}

	if err := Cfg.SaveTo(configFilePath); err != nil {
		log.Printf("error when saving to %s", configFilePath)
		panic(err)
	}
}

func mapTo(section string, v interface{}) {
	err := Cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

func reflectFrom(section string, v interface{}) {
	err := Cfg.Section(section).ReflectFrom(v)
	if err != nil {
		log.Fatalf("Cfg.ReflectFrom %s err: %v", section, err)
	}
}
