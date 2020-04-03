package flag

import (
	"github.com/jshzhj/ginframe/pkg/setting"
	"github.com/jshzhj/ginframe/pkg/shell"
	"flag"
	"fmt"
)

var (
	ConfigPath *string
	Reload *string
	Stop *string
)

func Setup() {
	ConfigPath = flag.String("conf", "conf/app.ini", "请输入配置文件路径")
	Reload     = flag.String("reload", "false", "服务热重启")
	Stop       = flag.String("stop", "false", "服务热关停")
	flag.Parse()
	if *Reload == "true" {
		shell.ExecShell(fmt.Sprintf("kill -1 %v", setting.Pid))
	}
	if *Stop == "true" {
		shell.ExecShell(fmt.Sprintf("kill -15 %v", setting.Pid))
	}
}
