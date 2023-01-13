package in

import (
	"github.com/tigerinus/xin/codegen"
	"github.com/tigerinus/xin/model"
)

func ActionTypeAdapter(actionType codegen.ActionType) model.ActionType {
	propertyTypeList := make([]model.PropertyType, 0)
	for _, propertyType := range actionType.PropertyTypeList {
		propertyTypeList = append(propertyTypeList, PropertyTypeAdapter(propertyType))
	}

	return model.ActionType{
		SourceID:         actionType.SourceID,
		Name:             actionType.Name,
		PropertyTypeList: propertyTypeList,
	}
}
