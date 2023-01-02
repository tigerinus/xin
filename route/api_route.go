package route

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/tigerinus/xin/codegen"
	"github.com/tigerinus/xin/service"
)

type APIRoute struct {
	services *service.Services
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewAPIRoute(services *service.Services) codegen.ServerInterface {
	return &APIRoute{
		services: services,
	}
}
