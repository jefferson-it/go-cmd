package main

import (
	"os"
	"os/exec"
	"runtime"

	"sys.cmd/utils"
)

func main() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("powershell", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()

		if err != nil {
			utils.Error("A error has ocurred on clear terminal: ", err.Error())
		}
		return
	}

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		utils.Error("A error has ocurred on clear terminal: ", err.Error())
	}
}
