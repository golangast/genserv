/*
Copyright © 2023 zendrulat@gmail.com
*/
package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/golangast/genserv/cmd"
)

func main() {
	cmd.Execute()

	Server()
	if err := ShellBash("cd assets/db/rqlite/bin && chmod 755 ./rqlited && ./rqlited -auth config.json  ~/node.1 ", "trying to run database server bash command"); err != nil {
		fmt.Println(err)
	}
}

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func ShellBash(s, errmess string) error {

	out, errout, err := Shellout(s)
	if err != nil {
		return err
	}
	if out != "" {
		fmt.Println(out)
	}
	if errout != "" {

		fmt.Println(errout)
	}

	return nil

}
