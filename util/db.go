package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/aceld/zinx/zlog"
	_ "github.com/go-sql-driver/mysql"
)

type DBList struct {
	BookStoreDB DBInfo `json:"BookStoreDB"`
}

type DBInfo struct {
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	User     string `json:"User"`
	Password string `json:"Password"`
	DBName   string `json:"DBName"`
}

var bookStoreDB *sql.DB

func initDB() {
	pwd, err := os.Getwd()
	if err != nil {
		pwd = "."
	}
	dbConfPath := pwd + "/conf/db.json"
	zlog.Infof("DB|dbConfPath|%s", dbConfPath)
	data, err := ioutil.ReadFile(dbConfPath)
	if err != nil {
		panic(err)
	}

	//将json数据解析到struct中
	var dbList0 DBList
	err = json.Unmarshal(data, &dbList0)
	if err != nil {
		panic(err)
	}

	//
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbList0.BookStoreDB.User,
		dbList0.BookStoreDB.Password,
		dbList0.BookStoreDB.Host,
		dbList0.BookStoreDB.Port,
		dbList0.BookStoreDB.DBName,
	)
	zlog.Infof("DB|Start|%v|%s", dbList0.BookStoreDB, dsn)

	bookStoreDB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = bookStoreDB.Ping()
	if err != nil {
		panic(err)
	}

	// 建议将此方法参数传递在五分钟之内，该设置对负载平衡和更改系统变量也很有帮助
	bookStoreDB.SetConnMaxLifetime(240 * time.Second)
	// 设置数据库最大连接数
	bookStoreDB.SetMaxOpenConns(20)
	// 设置上数据库最大闲置连接数
	bookStoreDB.SetMaxIdleConns(10)
}

func GetBookInfoDB() *sql.DB {
	return bookStoreDB
}
