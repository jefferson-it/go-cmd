package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Compile(target string, platform string, arch string) {
	// Exec compilation
	cmd := exec.Command("go", "build", target)

	// Add Env
	cmd.Env = os.Environ()

	// Add platform on GOOS env, if te uppercase, convert to lowercase
	if platform != "AUTO" {
		cmd.Env = append(cmd.Env, fmt.Sprintf("GOOS=%s", strings.ToLower(platform)))
	}
	// Add arch on GOARCH env, if te uppercase, convert to lowercase
	if arch != "AUTO" {
		cmd.Env = append(cmd.Env, fmt.Sprintf("GOARCH=%s", strings.ToLower(arch)))
	}

	// run compile command
	err := cmd.Run()

	if err != nil {
		fmt.Printf("A erros has ocurred on compile: %s.\n", err.Error())

		Exit()
	}

	// The command move bin file to folder /bin/
	move(target, platform)
}
