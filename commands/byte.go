package main

import (
	"fmt"
	"os"
)

func main() {
	text := ""

	if len(os.Args) > 1 {
		for i, data := range os.Args {
			if i == 0 {
				continue
			}
			if i == 1 {
				text = fmt.Sprintf("%s %v", text, data)
				continue
			}
			text = fmt.Sprintf(" %s %v", text, data)
		}
		fmt.Println([]byte(text))
		return
	}

	fmt.Println(`Require value for convert to byte:
'byte <...args>'
byte "Hello, World" 
output: [72 101 108 108 111 44 32 87 111 114 108 100]`)
}
