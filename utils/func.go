package utils

import (
	"io/ioutil"
	"os"
)

func ReadFile(filename string, callbackErr func(err error)) string {
	res, err := ioutil.ReadFile(filename)

	if err != nil {
		callbackErr(err)
	}

	return string(res)
}

func CreateFile(filename string, content ...string) {
	txt := "// This file is writing for file(Go Commands), github.com/jefferson-developer-it/go-cmd"

	if len(content) > 0 {
		txt = content[0]
	}

	err := ioutil.WriteFile(filename, []byte(txt), os.ModePerm)

	if err != nil {
		Error("Erro on create your file!\n", err.Error())
	}
}
