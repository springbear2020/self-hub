package initialize

import (
	"github.com/springbear2020/self-hub/server/global"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
