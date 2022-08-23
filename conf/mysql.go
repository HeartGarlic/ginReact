package conf

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Charset  string
}

var Engine *xorm.Engine

func InitMysql() {
	MysqlConf := Mysql{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "luo1010.",
		Database: "gin_react_blog",
		Charset:  "utf8",
	}
	// 启动Orm 引擎
	var err error
	Engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", MysqlConf.User, MysqlConf.Password, MysqlConf.Host, MysqlConf.Database, MysqlConf.Charset))
	if err != nil {
		panic(err)
	}
}
