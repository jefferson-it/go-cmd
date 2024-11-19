package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Exit() {
	os.Exit(0)
}

/*
 Read local path on terminal
*/
func HerePath() string {
	res, err := exec.Command("pwd").Output()

	if err != nil {
		Error("A erros has ocurred: %s.\n", err.Error())

		Exit()
	}

	return strings.Trim(string(res), "\n")
}

func CreateBin() {
	_, err := ioutil.ReadDir("./bin/")

	if err != nil {
		if err = os.Mkdir("./bin/", os.ModePerm); err != nil {
			Error("Error on create \"bin\" folder.", err.Error())
			Exit()
		}
		Log("The \"bin\" folder has create.")
	}

	Log("The \"bin\" folder has existent.")
}

func move(target string, platform string) {
	path_from := fmt.Sprintf(".%s", strings.Split(strings.Split(target, "/commands")[1], ".")[0])

	if strings.Contains(platform, "win") || runtime.GOOS == "windows" {
		path_from = fmt.Sprintf("%s.exe", path_from)
	}

	if runtime.GOOS == "windows" {

		cmd := exec.Command("powershell.exe", "mv", path_from, "./bin/")
		err := cmd.Run()

		if err != nil {
			Error(fmt.Sprintf("A erros has ocurred on move: %s.\n", path_from), err.Error())

			Exit()
		}

		Log(fmt.Sprintf("The file %s has compiled and moved to bin folder.", strings.Split(path_from, "./")[1]))

		return
	}

	cmd := exec.Command("mv", path_from, "./bin/")
	err := cmd.Run()

	if err != nil {
		Error("A erros has ocurred on move: %s.\n", err.Error())

		Exit()
	}

	Log(fmt.Sprintf("The file %s has compiled and moved to bin folder.", strings.Split(path_from, "./")[1]))
}
