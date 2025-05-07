package test

import (
	"bytes"
	"fmt"
	"os/exec"
)

// RunNodeJSCode 运行 Node.js 代码
func RunNodeJSCode(code string) (string, error) {
	// 创建一个执行 Node.js 的命令
	cmd := exec.Command("node", "-e", code)

	// 创建一个缓冲区来存储命令的输出
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// 执行命令
	err := cmd.Run()
	if err != nil {
		// 如果执行出错，返回错误信息
		return "", fmt.Errorf("执行 Node.js 代码时出错: %v, 错误输出: %s", err, stderr.String())
	}

	// 返回命令的输出结果
	return out.String(), nil
}
