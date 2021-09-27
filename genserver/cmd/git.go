package cmd

import (
	"log"
	"os"
	"os/exec"
)

//find . -name "*.back.txt"  | xargs rm -f

func GitAdd(projectPath string) {
	err := os.Chdir(projectPath)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("git", "add", ".")
	stdout, err := cmd.StdoutPipe()
	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		log.Fatal(err)
	}
	defer stdout.Close()                // 保证关闭输出流
	if err := cmd.Start(); err != nil { // 运行命令
		log.Fatal(err)
	}
}
