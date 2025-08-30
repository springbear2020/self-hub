package response

import (
	"github.com/springbear2020/self-hub/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
