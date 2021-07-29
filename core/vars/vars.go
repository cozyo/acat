package vars

import (
	"github.com/acat/core/config/setting"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"

	"gitee.com/kelvins-io/common/log"
)

var (
	App                    *WebApplication
	DBEngineGORM           *gorm.DB
	MysqlSettingMicroGM    *setting.MysqlSettings
	RedisSettingMicroGM    *setting.RedisSettings
	VerifyCodeSetting      *VerifyCodeSettingS
	LoggerSetting          *setting.LoggerSettings
	ErrorLogger            log.LoggerContextIface
	AccessLogger           log.LoggerContextIface
	RedisPoolEngine        *redis.Pool
	ServerSetting          *setting.ServerSettings
)

