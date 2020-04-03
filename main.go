package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/jshzhj/ginframe/pkg/file"
	"github.com/jshzhj/ginframe/pkg/flag"
	"github.com/jshzhj/ginframe/pkg/logger"
	"github.com/jshzhj/ginframe/pkg/mysql"
	"github.com/jshzhj/ginframe/pkg/redis"
	"github.com/jshzhj/ginframe/pkg/setting"
	"github.com/jshzhj/ginframe/routers"
	"log"
	"strconv"
	"syscall"
	"time"
)

func init(){
	setting.Setup("./conf/app.ini")
	mysql.Setup()
	redis.Setup()
	logger.Setup()
	flag.Setup()
}

func main() {
	gin.SetMode(setting.HttpServerSetting.RunMode)
	endless.DefaultReadTimeOut = setting.HttpServerSetting.ReadTimeout * time.Second
	endless.DefaultWriteTimeOut = setting.HttpServerSetting.WriteTimeout * time.Second
	endless.DefaultHammerTime = setting.HttpServerSetting.CloseWaitTime * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20

	router := routers.InitRouter()
	server := endless.NewServer(":"+setting.HttpServerSetting.Port, router)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
		//创建pid文件
		pid := syscall.Getpid()
		if err := file.CreateFile("./runtime/pid.txt");err!=nil{
			log.Fatal("创建pid缓存文件失败:",err)
		}
		if err := file.WriteFile("./runtime/pid.txt","",strconv.Itoa(pid));err !=nil{
			log.Fatal("写pid缓存文件失败:",err)
		}
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
