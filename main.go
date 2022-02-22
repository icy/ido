/*
  Purpose : Execute a command provided as arguments,
            ... but aks user to confirm if there is some special string
            from the arguments (live, delete, rm, remove,...)
  Author  : Ky-Anh Huynh
  Github  : https://github.com/icy/ido
  Date    : 2022-02-22
  License : MIT
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"syscall"
)

var regExpDefault = regexp.MustCompile("(?i)(del|delete|remove|rm|live)")

func warnf(format string, a ...interface{}) {
	fmt.Fprint(os.Stderr, ":: ", fmt.Sprintf(format, a...))
}

func main() {
	flag.Parse()
	cmdArgs := flag.Args()
	if len(cmdArgs) < 1 {
		warnf("Please specify a command.\n")
		os.Exit(1)
	}
	icount := 0
	for _, arg := range cmdArgs {
		if regExpDefault.FindStringIndex(arg) != nil {
			icount += 1
			warnf("Found risky pattern: %s\n", arg)
		}
	}
	if icount > 0 {
		warnf("Please type YES and enter to continue: ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		if input.Text() == "YES" {
			warnf("Going to execute your command... Best luck.\n")
		} else {
			warnf("Thanks, you may have saved your system.\n")
			os.Exit(1)
		}
	}

	execPath := cmdArgs[0]
	execPathResolved, err := exec.LookPath(execPath)
	if err != nil {
		panic(err)
	}

	syscall.Exec(execPathResolved, cmdArgs, syscall.Environ())
}
