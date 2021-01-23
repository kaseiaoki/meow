package executeCmd

import (
	"fmt"
	"os/exec"
	"runtime"
)

func Out(arg string) (string, error){
	var out []byte
	var err error

	if runtime.GOOS == "windows" {
		out, err = exec.Command("cmd", "/C", arg).Output()
	} else {
		out, err = exec.Command("bash", "-c", "<your command>").Output()
	}

	if err != nil {
		fmt.Println(err);
		return "", err
	}
	return string(out), nil
}
