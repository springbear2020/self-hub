package response

import "github.com/springbear2020/self-hub/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
