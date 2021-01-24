package executeCmd

import (
	"fmt"
	"io"
	"os/exec"
	"runtime"
	// "os"
)

func Out(arg string) (string, error) {
	var out []byte
	var err error
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", arg)
	} else {
		cmd = exec.Command("bash", "-c", "<your command>")
	}
	out, err = cmd.Output()

	return string(out), err
}

func StdIO(arg string, input string) (string, error) {
	var out []byte
	var err error
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", arg)
	} else {
		cmd = exec.Command("bash", "-c", "<your command>")
	}

	stdin, _ := cmd.StdinPipe()
	io.WriteString(stdin, input)
	stdin.Close()
	out, err = cmd.Output()

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(out), nil
}
