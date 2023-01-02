package out

import (
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/utils"
	"github.com/tigerinus/xin/codegen"
	"github.com/tigerinus/xin/model"
)

func ActionAdapter(action model.Action) codegen.Action {
	return codegen.Action{
		SourceID:   action.SourceID,
		Name:       action.Name,
		Properties: action.Properties,
		Timestamp:  utils.Ptr(time.Unix(action.Timestamp, 0)),
	}
}
