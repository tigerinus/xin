package out

import (
	"github.com/tigerinus/xin/codegen"
	"github.com/tigerinus/xin/model"
)

func ActionTypeAdapter(actionType model.ActionType) codegen.ActionType {
	propertyTypeList := make([]codegen.PropertyType, 0)
	for _, propertyType := range actionType.PropertyTypeList {
		propertyTypeList = append(propertyTypeList, PropertyTypeAdapter(propertyType))
	}

	return codegen.ActionType{
		SourceID:         actionType.SourceID,
		Name:             actionType.Name,
		PropertyTypeList: propertyTypeList,
	}
}
