package utils

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

// 执行 linux shell command
func ExecLinuxShell(s string) (string, error) {
	//函数返回一个io.Writer类型的*Cmd
	cmd := exec.Command("/bin/bash", "-c", s)

	//通过bytes.Buffer将byte类型转化为string类型
	var result bytes.Buffer
	cmd.Stdout = &result

	//Run执行cmd包含的命令，并阻塞直至完成
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return result.String(), err
}
func Reload() error {
	if runtime.GOOS == "windows" {
		return errors.New("系统不支持")
	}
	pid := os.Getpid()
	cmd := exec.Command("kill", "-1", strconv.Itoa(pid))
	return cmd.Run()
}
