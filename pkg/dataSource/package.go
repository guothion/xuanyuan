package dataSource

import (
	"github.com/guothion/xuanyuan/pkg/config"
	"github.com/guothion/xuanyuan/pkg/dataSource/mysql"
)

func Init() {
	if config.Config.DataSource.Type == "mysql" {
		mysql.Init()
	} else if config.Config.DataSource.Type == "mongo" {

	} else if config.Config.DataSource.Type == "postgres" {

	} else if config.Config.DataSource.Type == "redis" {

	} else if config.Config.DataSource.Type == "sqlite" {

	}
}
