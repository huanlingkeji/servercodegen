package cmd

import (
	"os"
	"os/exec"
)

// find . -name "*.back.txt"  | xargs rm -f

// GitAdd GitAdd
func GitAdd(projectPath string) error {
	err := os.Chdir(projectPath)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "add", ".")
	stdout, err := cmd.StdoutPipe()
	if err != nil { // 获 取输出对象，可以从该对象中读取输出结果
		return err
	}
	defer stdout.Close()                // 保证关闭输出流
	if err := cmd.Start(); err != nil { // 运行命令
		return err
	}
	return nil
}
