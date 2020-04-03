package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jshzhj/ginframe/pkg/setting"
	//"github.com/xormplus/core"
	"github.com/xormplus/xorm"
	"log"
)

var DB *xorm.Engine

func Setup(){
	var err error

	dsn := setting.MysqlSetting.User + ":"
	dsn += setting.MysqlSetting.Pwd + "@"
	dsn += "(" + setting.MysqlSetting.Ip + ":"
	dsn += setting.MysqlSetting.Port + ")/"
	dsn += setting.MysqlSetting.Database + "?charset=utf8"

	DB, err = xorm.NewEngine("mysql", dsn)

	if err = DB.Ping(); err != nil {
		log.Fatalf("mysql.Setup,mysql连接失败:%v", err.Error())
	}

	DB.ShowSQL(setting.MysqlSetting.ShowSQL)                     //打印sql
	//DB.Logger().SetLevel(core.LOG_DEBUG) //打印调试
	DB.SetMaxOpenConns(setting.MysqlSetting.SetMaxOpenConns)
	DB.SetMaxIdleConns(setting.MysqlSetting.SetMaxIdleConns)//mysql连接池最大连接数

	return
}

func Close() {
	defer DB.Close()
}