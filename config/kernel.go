package config

import (
	"github.com/acat/core/config"
	"log"
	"strings"
)

func MapConfig(section string, v interface{})  {
	if strings.HasPrefix(section, "web-") {
		log.Fatalf("[err] section name can't have web- perfix")
	}
	config.MapConfig(section, v)
}