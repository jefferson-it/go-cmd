package main

import (
	"fmt"
	"os"

	"sys.cmd/utils"
)

func main() {
	var flag, filename, content string

	flag = "auto"

	if len(os.Args) > 1 {
		if os.Args[1] == "-r" || os.Args[1] == "-w" || os.Args[1] == "-ow" || os.Args[1] == "-h" {
			flag = os.Args[1]
		} else {
			filename = os.Args[1]
		}

		if len(os.Args) > 2 && filename == "" {
			filename = os.Args[2]
		} else if len(os.Args) > 2 && filename != "" {
			content = os.Args[2]
		}

		if len(os.Args) > 3 && content == "" {
			content = os.Args[3]
		}
	}

	if filename == "" && flag != "-h" {
		utils.Error("Please, select filename! 'file <flag> <filename> <content>'\nType 'file -h' for help list.")
		utils.Exit()
	}
	if flag == "-h" {

		fmt.Println(`
Thank'you to use my command!
> The binary 'file' work whit file, reading and writing file.
use 'file <flag> <filename> <content>'

flag:
	-r = read a file
	-w = write a new file
	-ow = overwrite a existent file 
	-h show this screen

If first argument was not a any of the list, then this value was value of filename.
And flag was work as auto, first exec read(if file exist), else write.`)

	} else if flag == "auto" {
		res := utils.ReadFile(filename, func(err error) {
			utils.CreateFile(filename, content)
			utils.Exit()
		})

		fmt.Println(res)
	} else if flag == "-r" {
		res := utils.ReadFile(filename, func(err error) {
			utils.Error(err.Error())
			utils.Exit()
		})
		fmt.Println(res)
	} else if flag == "-w" {
		res := utils.ReadFile(filename, func(err error) {
			utils.CreateFile(filename, content)
			utils.Exit()
		})

		if res != "" {
			utils.Error("This file exist!!\n", " Use 'file -ow <filename> <content>' for overwriting!")
			utils.Exit()
		}
	} else if flag == "-ow" {
		utils.CreateFile(filename, content)
	}
}
