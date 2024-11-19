package utils

import "fmt"

func Log(args ...interface{}) {
	msg := fmt.Sprint(ColorGreen)

	for _, arg := range args {
		msg = fmt.Sprintf(" %s %v ", msg, arg)
	}

	fmt.Println(msg)
}

func Error(args ...interface{}) {
	msg := fmt.Sprint(ColorRed)

	for _, arg := range args {
		msg = fmt.Sprintf(" %s %v ", msg, arg)
	}

	fmt.Println(msg)
}

const ColorReset = "\033[0m"
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue = "\033[34m"
const ColorPurple = "\033[35m"
const ColorCyan = "\033[36m"
const ColorWhite = "\033[37m"
