package in

import (
	"github.com/tigerinus/xin/codegen"
	"github.com/tigerinus/xin/model"
)

func PropertyTypeAdapter(propertyType codegen.PropertyType) model.PropertyType {
	return model.PropertyType{
		Name: propertyType.Name,
	}
}
