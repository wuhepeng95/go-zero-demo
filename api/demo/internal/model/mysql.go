package model

import (
	"demo/internal/config"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func GetMysqlConn(conf config.Config) (con sqlx.SqlConn) {
	var InformationSchemaDsn = fmt.Sprintf("%v:%v@tcp(%v:3306)/%v",
		conf.MysqlUsername, conf.MysqlPassword, conf.MysqlHost, conf.MysqlDatabase)
	return sqlx.NewMysql(InformationSchemaDsn)
}
