package startup

import (
	"github.com/acat/config"
	"github.com/acat/core/config/setting"
	"github.com/acat/core/vars"
	"log"
)


const (
	SectionMysqlMicroGM = "micro-mall-mysql"
	SectionRedisMicroGM = "micro-mall-redis"
	SectionVerifyCode     = "micro-mall-verify_code"
)

func LoadConfig() error {

	// 外部MySQL数据源
	log.Printf("[info] Load default config %s", SectionMysqlMicroGM)
	vars.MysqlSettingMicroGM = new(setting.MysqlSettings)
	config.MapConfig(SectionMysqlMicroGM, vars.MysqlSettingMicroGM)

	// 加载外部Redis数据源
	log.Printf("[info] Load default config %s", SectionRedisMicroGM)
	vars.RedisSettingMicroGM = new(setting.RedisSettings)
	config.MapConfig(SectionRedisMicroGM, vars.RedisSettingMicroGM)

	//加载G2Cache二级缓存配置


	// 加载验证码配置
	log.Printf("[info] Load default config %s", SectionVerifyCode)
	vars.VerifyCodeSetting = new(vars.VerifyCodeSettingS)
	config.MapConfig(SectionVerifyCode, vars.VerifyCodeSetting)
	return nil
}