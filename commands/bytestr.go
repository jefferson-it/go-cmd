package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"sys.cmd/utils"
)

func main() {
	text := ""

	if len(os.Args) > 1 {
		var byte_res []byte
		for i, data := range os.Args {
			if i == 0 {
				continue
			} else if i == 1 && strings.Contains(data, "[") {
				text = fmt.Sprintf("%s %v", text, data)
				continue
			} else if i == 1 && !strings.Contains(data, "[") {
				text = fmt.Sprintf("%s %s %v", text, "[ ", data)
				continue
			}

			text = fmt.Sprintf(" %s %v", text, data)
		}

		if !strings.Contains(text, "]") {
			text = fmt.Sprintf("%s %s", text, " ]")
		}

		err := json.Unmarshal([]byte(text), &byte_res)

		if err != nil {
			utils.Error(err.Error())
		}

		fmt.Println(string(byte_res))
		return
	}

	fmt.Println(`Require value for convert to string:
'bytestr <...args>'
bytestr [ 72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100 ] 
output: "Hello, World"`)
}
