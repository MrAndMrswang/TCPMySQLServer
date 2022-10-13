package dao

import (
	"TCPMySQLServer/util"
	"TCPMySQLServer/vo"

	"github.com/aceld/zinx/zlog"
)

func Query(id0 int64) []vo.Book {
	sqlStr := "select id, name, remark, created_by from bookInfo where id > ?"

	db0 := util.GetBookInfoDB()
	rows, err := db0.Query(sqlStr, id0)
	if err != nil {
		zlog.Errorf("Query|%v", err)
		return nil
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	var bookInfoList0 []vo.Book
	for rows.Next() {
		var info0 vo.Book
		err := rows.Scan(
			&info0.Id,
			&info0.Name,
			&info0.Remark,
			&info0.CreatedBy,
		)
		if err != nil {
			zlog.Errorf("Select|%v", err)
			return nil
		}

		bookInfoList0 = append(bookInfoList0, info0)
	}
	zlog.Infof("Select|bookInfoList0=%+v", bookInfoList0)
	return bookInfoList0
}
