package config

import (
	"bytes"
	"fmt"
	"gitee.com/kelvins-io/common/env"
	"github.com/acat/core/config/setting"
	"github.com/jinzhu/gorm"
	"net/url"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLGORMEngine(settings *setting.MysqlSettings) (*gorm.DB, error)  {
	if settings == nil {
		return nil, fmt.Errorf("Mysql setting is nil")
	}
	if settings.UserName == "" {
		return nil, fmt.Errorf("lack of mysql setting UserName")
	}
	if settings.Password == "" {
		return nil, fmt.Errorf("lack of mysql setting Password")
	}
	if settings.Host == "" {
		return nil, fmt.Errorf("lack of mysql setting Host")
	}
	if settings.DataBase == "" {
		return nil, fmt.Errorf("lack of mysql setting DataBase Name")
	}
	if settings.Charset == "" {
		return nil, fmt.Errorf("lack of mysql setting Charset")
	}
	if settings.PoolNum == 0 {
		return nil, fmt.Errorf("lack of mysql setting PoolNum")
	}

	var buf bytes.Buffer
	buf.WriteString(settings.UserName)
	buf.WriteString(":")
	buf.WriteString(settings.Password)
	buf.WriteString("@tcp(")
	buf.WriteString(settings.Host)
	buf.WriteString(":")
	buf.WriteString(strconv.Itoa(settings.Port))
	buf.WriteString(")/")
	buf.WriteString(settings.DataBase)
	buf.WriteString("?charset=")
	buf.WriteString(settings.Charset)
	buf.WriteString("&parseTime=" + strconv.FormatBool(settings.ParseTime))
	buf.WriteString("&multiStatements=" + strconv.FormatBool(settings.MultiStatements))
	if settings.Loc == "" {
		buf.WriteString("&loc=Local")
	} else {
		buf.WriteString("&loc=" + url.QueryEscape(settings.Loc))
	}
	db, err := gorm.Open("mysql", buf.String())
	if err != nil{
		return nil, err
	}
	if env.IsDevMode() {
		db.LogMode(true)
	}

	if settings.ConnMaxLifeSecond <= 0 {
		settings.ConnMaxLifeSecond = 30
	}
	db.DB().SetConnMaxLifetime(time.Duration(settings.ConnMaxLifeSecond) * time.Second)

	if settings.MaxIdleConns <= 0 {
		settings.MaxIdleConns = 10
	}
	db.DB().SetMaxIdleConns(settings.MaxIdleConns)

	if settings.PoolNum <= 0 {
		settings.PoolNum = 10
	}
	db.DB().SetMaxOpenConns(settings.PoolNum)

	return db, nil
}