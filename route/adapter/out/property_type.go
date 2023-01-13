package out

import (
	"github.com/tigerinus/xin/codegen"
	"github.com/tigerinus/xin/model"
)

func PropertyTypeAdapter(propertyType model.PropertyType) codegen.PropertyType {
	return codegen.PropertyType{
		Name: propertyType.Name,
	}
}
