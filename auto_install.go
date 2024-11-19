package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"sys.cmd/utils"
)

func main() {
	/*
	*	Here define platform and arch to compile
	* 	WIN - windows
	* 	LINUX - linux
	* 	DARWIN - MacOS
	 */
	platform := "AUTO"
	arch := "AUTO"

	// Get info ./commands
	commands_folder_byte, err := ioutil.ReadDir("./commands")

	if err != nil {
		fmt.Printf("A erros has ocurred: %s.\n", err.Error())

		utils.Exit()
	}

	// Read arguments line, for append data on platform/arch
	if len(os.Args) > 1 {
		platform = os.Args[1]
		if len(os.Args) > 2 {
			arch = os.Args[2]
		}
	}

	// Create bin folder(if not exist)
	utils.CreateBin()

	// Range data of "./commands" for compile files
	for _, data := range commands_folder_byte {
		utils.Compile(fmt.Sprintf("./commands/%s", data.Name()), platform, arch)
	}

	utils.Exit()
}
