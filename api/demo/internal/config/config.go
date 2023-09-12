package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	MysqlHost     string `json:"env=MYSQL_HOST"`
	MysqlUsername string `json:"env=MYSQL_USERNAME"`
	MysqlPassword string `json:"env=MYSQL_PASSWORD"`
	MysqlDatabase string `json:"env=MYSQL_DATABASE"`
}
