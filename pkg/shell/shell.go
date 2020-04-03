package shell

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecShell(arg string) string {
	var (
		output []byte
		err    error
	)
	fmt.Println(arg)
	cmd := exec.Command("/bin/bash", "-c", arg)

	if output, err = cmd.CombinedOutput(); err != nil {
		log.Fatalf("flag.ExecSgell:执行shell命令失败: %v",err)
	}
	log.Println("命令执行成功")
	return string(output)
}
