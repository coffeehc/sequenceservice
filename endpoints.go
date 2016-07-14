package sequenceservice

import (
	"fmt"

	"github.com/coffeehc/microserviceboot/base"
	"github.com/coffeehc/web"
)

var (
	PATHPARAM_SQUENCE   = "sequenceId"
	PATHPARAM_TIMESTEMP = "timestemp"
)

var (
	POST_SEQUENCE   = base.EndPointMeta{Path: "/api/v1/sequences", Method: web.POST, Description: "创建一个新的 sequence"}
	GET_SEQUENCE    = base.EndPointMeta{Path: fmt.Sprintf("/api/v1/sequence/{%s}", PATHPARAM_SQUENCE), Method: web.GET, Description: "获取 Sequence 元信息"}
	GET_MINSEQUENCE = base.EndPointMeta{Path: fmt.Sprintf("/api/v1/sequences/min/{%s}", PATHPARAM_TIMESTEMP), Method: web.GET, Description: "获取给定时间戳里最小的 sequenceId"}
)
