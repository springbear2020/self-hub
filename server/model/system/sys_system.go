package system

import (
	"github.com/springbear2020/self-hub/server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
