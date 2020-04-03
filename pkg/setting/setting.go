package setting

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)
type App struct {}
var AppSetting = &App{}

type Log struct {
	LogLevel    int
	LogSavePath string
	MaxSize     int
	MaxBackups  int
	MaxAge      int
	Compress    bool
	IsConsole   int
}

var LogSetting = &Log{
	LogLevel:    0,
	LogSavePath: "./runtime/logs/gin.log",
	MaxSize:     100,
	MaxBackups:  1000,
	MaxAge:      30,
	Compress:    true,
	IsConsole:   1,
}

type HttpServer struct {
	RunMode       string
	Port          string
	CloseWaitTime time.Duration
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
}

var HttpServerSetting = &HttpServer{
	RunMode: "debug",
}

type Mysql struct {
	User            string
	Pwd             string
	Ip              string
	Port            string
	Database        string
	ShowSQL         bool
	SetLevel        int
	SetMaxOpenConns int
	SetMaxIdleConns int
}

var MysqlSetting = &Mysql{}

type Redis struct{}

var RedisSetting = &Redis{}

var (
	cfg *ini.File
	Pid int //主进程pid
)

func Setup(confPath string) {
	var err error

	//加载配置文件
	if cfg, err = ini.Load(confPath); err != nil {
		log.Fatalf("setting.Setup,加载配置文件:'%v'失败: %v", confPath, err)
	}

	//映射到结构体
	mapTo("App", AppSetting)
	mapTo("Log", RedisSetting)
	mapTo("HttpServer", HttpServerSetting)
	mapTo("Mysql", MysqlSetting)
	mapTo("Redis", RedisSetting)


	return
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
