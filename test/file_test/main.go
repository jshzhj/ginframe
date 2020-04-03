package main

import (
	"fmt"
	"os"
)

func main(){
	var (
		pwd string
		err error
	)
	//创建日志目录
	if pwd,err = os.Getwd();err !=nil{
		fmt.Println("获取当前路径失败")
	}
	fmt.Println(pwd)
	os.Exit(1)
}
