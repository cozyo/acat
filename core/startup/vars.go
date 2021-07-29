package startup

import (
	"gitee.com/kelvins-io/common/log"
	"github.com/acat/config"
	"github.com/acat/core/vars"
)

func SetupVars() error {
	var err error

	vars.ErrorLogger, err = log.GetErrLogger("err")
	if err != nil {
		return err
	}

	vars.AccessLogger, err = log.GetAccessLogger("access")
	if err != nil {
		return err
	}

	if vars.MysqlSettingMicroGM != nil && vars.MysqlSettingMicroGM.Host != ""{
		vars.DBEngineGORM, err = config.NewMySQLGORMEngine(vars.MysqlSettingMicroGM)
		if err != nil {
			return err
		}
	}

	if vars.RedisSettingMicroGM != nil && vars.RedisSettingMicroGM.Host != ""{
		vars.RedisPoolEngine, err = config.NewRedis(vars.RedisSettingMicroGM)
		if err != nil {
			return err
		}
	}

	return nil
}

func SetStopFunc() (err error) {
	if vars.RedisPoolEngine != nil {
		if err := vars.RedisPoolEngine.Close(); err != nil {
			return err
		}
	}

	return err
}