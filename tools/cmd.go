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
	stdErr, _ := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
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
	errSS := []string{}
	errReader := bufio.NewReader(stdErr)
	for {
		line, err2 := errReader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Print(line)
		errSS = append(errSS, line)
	}
	cmd.Wait()
	if len(errSS) > 0 {
		return true, append(ss, errSS...)
	}
	return true, ss
}
