package global

import (
	"github.com/jinzhu/gorm"
	"on_one_heels_go/pkg/settings"
)

var (
	ServerSetting   *settings.ServerSetting
	DatabaseSetting *settings.DataBaseSetting
	DBEngine        *gorm.DB
	UniSwapSetting  *settings.UniSwapSetting
)
