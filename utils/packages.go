package utils

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteCommand(command string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	_ = cmd.Wait()
}

func CheckPackage(arg string) bool {
	command := exec.Command("apt", "-qq", "list", arg)
	stdout, err := command.Output()
	if err != nil {
		ShowMessage("Erreur !", true)
	} else if strings.Contains(string(stdout), "installed") {
		return true
	}
	return false
}
