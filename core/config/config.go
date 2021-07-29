package config

import (
	"flag"
	"github.com/acat/core/config/setting"
	"github.com/acat/core/vars"
	"github.com/go-ini/ini"
	"log"
)

const (
	// ConfFileName defines config file name.
	ConfFileName = ".env"
	// SectionServer is a section name for grpc server.
	SectionServer = "web-server"
	// SectionLogger is a section name for logger.
	SectionLogger = "web-logger"
	// SectionMysql is a sectoin name for mysql.
	SectionMysql = "web-mysql"
	// SectionRedis is a section name for redis.
	SectionRedis = "web-redis"
)

var (
	cfg        *ini.File
	flagConf = flag.String("conf_file", "", "Set app config.")
)

func LoadDefaultConfig() error {
	// Setup cfg object
	flag.Parse()
	var err error
	var confFile = ConfFileName
	if *flagConf != "" {
		confFile = *flagConf
	}
	cfg, err = ini.Load(confFile)
	if err != nil {
		return err
	}

	// Setup default settings
	for _, sectionName := range cfg.SectionStrings() {
		if sectionName == SectionServer {
			log.Printf("[info] Load default config %s", sectionName)
			vars.ServerSetting = new(setting.ServerSettings)
			MapConfig(sectionName, vars.ServerSetting)
			continue
		}
		if sectionName == SectionLogger {
			log.Printf("[info] Load default config %s", sectionName)
			vars.LoggerSetting = new(setting.LoggerSettings)
			MapConfig(sectionName, vars.LoggerSetting)
			continue
		}
	}
	return nil
}


// MapConfig uses cfg to map config.
func MapConfig(section string, v interface{})  {
	sec, err := cfg.GetSection(section)
	if err != nil {
		log.Fatalf("[err] Fail to parse '%s': %v", section, err)
	}
	if err := sec.MapTo(v); err != nil {
		log.Fatalf("[err] %s section map to setting err； %v", section, err)
	}
}