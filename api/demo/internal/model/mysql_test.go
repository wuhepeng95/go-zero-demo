package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"testing"
)

func Test1(t *testing.T) {
	dsn := ""
	// 创建一个sqlx连接
	conn := sqlx.NewSqlConn("mysql", dsn)

	data, err := NewSlowLogModel(conn).FindOne(context.Background(), 283931)
	if err != nil {
		if err == ErrNotFound {
			// 处理数据不存在的情况
		}

	}
	fmt.Println(data)
}
func Test2(t *testing.T) {
	dsn := ""
	// 创建一个sqlx连接
	conn := sqlx.NewSqlConn("mysql", dsn)

	sqlId := "srm"
	startTime := int64(1694015759000)
	sortBy := "rows_examined desc" // 按 RowsExamined 字段排序
	limit := 100
	offset := 0

	dataList, err := NewSlowLogModel(conn).FindByQuery(context.Background(), sqlId, &startTime, nil, sortBy, limit, offset)
	if err != nil {
		// 处理查询错误
		fmt.Printf("查询失败:%v\n", err)
	} else {
		fmt.Printf("查询成功:%v\n", len(dataList))
	}

}
