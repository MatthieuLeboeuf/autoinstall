package utils

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
	"gopkg.in/ini.v1"
	"os"
	"os/exec"
	"strings"
)

func GoodBye() {
	fmt.Println("Merci d'avoir utilisé", Cyan("AutoUpdater").Bold(), "!")
	os.Exit(0)
}

func ShowMessage(message string, error bool) {
	fmt.Println(Cyan("AutoUpdater").Bold(), "-", message)
	if error {
		os.Exit(1)
	}
}

func ReadOSRelease() map[string]string {
	cfg, err := ini.Load("/etc/os-release")
	if err != nil {
		ShowMessage("Une erreur est survenue !", true)
	}
	ConfigParams := make(map[string]string)
	ConfigParams["ID"] = cfg.Section("").Key("ID").String()
	ConfigParams["VERSION_ID"] = cfg.Section("").Key("VERSION_ID").String()
	return ConfigParams
}

func GetOsName() string {
	command := exec.Command("lsb_release", "-sc")
	stdout, err := command.Output()
	if err != nil {
		ShowMessage("Erreur !", true)
	}
	return strings.ReplaceAll(string(stdout), "\n", "")
}

func GetHostname() string {
	command := exec.Command("hostname")
	stdout, err := command.Output()
	if err != nil {
		ShowMessage("Erreur !", true)
	}
	return strings.ReplaceAll(string(stdout), "\n", "")
}
