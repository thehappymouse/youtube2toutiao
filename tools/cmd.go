package tools

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func ExecCommand(commandName string, params []string) (r bool, ss []string) {
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Print(line)
		ss = append(ss, line)
	}
	cmd.Wait()
	return true, ss
}
