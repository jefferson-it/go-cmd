package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/tabwriter"

	"sys.cmd/utils"
)

func main() {
	args := os.Args[1:]
	flags := ""
	path := utils.HerePath()
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	for _, arg := range args {
		if strings.Contains(arg, "-") {
			flags = fmt.Sprintf("%s %v", flags, arg)
		} else {
			path = arg
		}
	}

	dir_info, err := ioutil.ReadDir(path)

	if err != nil {
		utils.Error("Erro on read a dir: ", err.Error())
		utils.Exit()
	}

	fmt.Fprintln(w, "Name\tType\tSize")
	fmt.Fprintln(w, "-----\t-----\t-----")

	if strings.Contains(flags, "-h") {
		fmt.Println("Use lj for view yours directories. Same ls.")
		utils.Exit()
	}

	for _, data := range dir_info {
		text := fmt.Sprintf(
			"%s\t%s\t%vkb",
			data.Name(),
			utils.OptStr(data.IsDir(), "dir", "file"),
			data.Size(),
		)

		fmt.Fprintln(w, text)
	}

	w.Flush()

}
