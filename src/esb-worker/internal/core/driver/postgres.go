package driver

import (
	config "esb-worker/configs"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
	"net/url"
)

type Postgres struct {}

func (p *Postgres) config(setting *config.Config) string {

	if setting.Driver.Host != "" {
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s&sslmode=%s", url.QueryEscape(setting.Driver.User), url.QueryEscape(setting.Driver.Pass),
			setting.Driver.Host, setting.Driver.Port, setting.Driver.Db, setting.Driver.Schema, setting.Driver.SslMode)

		if setting.Driver.Schema == "false" {
			connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", url.QueryEscape(setting.Driver.User), url.QueryEscape(setting.Driver.Pass),
				setting.Driver.Host, setting.Driver.Port, setting.Driver.Db, setting.Driver.SslMode)
		}

		return connStr
	}

	return ""
}

func (p *Postgres) Engine(setting *config.Config) (*gorm.DB, error) {
	connectString := p.config(setting)
	if connectString != "" {
		engine, err := gorm.Open("postgres", connectString)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		engine.DB().SetMaxIdleConns(0)
		//engine.DB().SetMaxOpenConns(_maxConn)
		engine.LogMode(true)
		return engine, err
	}
	return nil, nil
}

var adapter Postgres
var Connection *gorm.DB