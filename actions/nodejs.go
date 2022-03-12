package actions

import (
	"encoding/json"
	"github.com/MatthieuLeboeuf/autoinstall/utils"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

func selectVersion() string {
	type version struct {
		Start string `json:"start"`
		End   string `json:"end"`
	}
	type versions map[string]version

	res, err := http.Get("https://raw.githubusercontent.com/nodejs/Release/main/schedule.json")
	body, err := ioutil.ReadAll(res.Body)
	var data versions
	json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}

	var nodeVersions []string

	for k, v := range data {
		l := "2006-01-02"
		start, _ := time.Parse(l, v.Start)
		end, _ := time.Parse(l, v.End)
		if time.Now().After(start) && time.Now().Before(end) {
			nodeVersions = append(nodeVersions, strings.ReplaceAll(k, "v", ""))
		}
	}

	sort.SliceStable(nodeVersions, func(i, j int) bool {
		return nodeVersions[i] < nodeVersions[j]
	})

	prompt := promptui.Select{
		Label:    "Quelle version voulez-vous",
		Items:    nodeVersions,
		HideHelp: true,
	}
	_, result, err := prompt.Run()
	if err != nil {
		utils.ShowMessage("Une erreur est survenue !", true)
	}
	return result
}

func yarnInstall() {
	if utils.CheckPackage("yarn") {
		utils.ShowMessage("Yarn est déja installé sur le système", false)
		prompt := promptui.Select{
			Label:    "Voulez-vous le remplacer",
			Items:    []string{"Oui", "Non"},
			HideHelp: true,
		}
		_, result, err := prompt.Run()
		if err != nil {
			utils.ShowMessage("Une erreur est survenue !", true)
		}
		if result == "Non" {
			utils.GoodBye()
		}
		utils.ExecuteCommand("apt-get remove --purge yarn -y")
		utils.ExecuteCommand("rm /usr/share/keyrings/yarnkey.gpg /etc/apt/sources.list.d/yarn.list")
	}
	utils.ExecuteCommand("curl -s https://dl.yarnpkg.com/debian/pubkey.gpg | gpg --dearmor | tee /usr/share/keyrings/yarnkey.gpg >/dev/null")
	utils.ExecuteCommand("echo 'deb [signed-by=/usr/share/keyrings/yarnkey.gpg] https://dl.yarnpkg.com/debian stable main' > /etc/apt/sources.list.d/yarn.list")
	utils.ExecuteCommand("apt-get update -y")
	utils.ExecuteCommand("apt-get install yarn -y")
}

func NodeJsInstall() {
	gpgFile := "/usr/share/keyrings/nodesource.gpg"

	if utils.CheckPackage("nodejs") {
		utils.ShowMessage("NodeJs est déja installé sur le système", false)
		prompt := promptui.Select{
			Label:    "Voulez-vous le remplacer",
			Items:    []string{"Oui", "Non"},
			HideHelp: true,
		}
		_, result, err := prompt.Run()
		if err != nil {
			utils.ShowMessage("Une erreur est survenue !", true)
		}
		if result == "Non" {
			utils.GoodBye()
		}
		utils.ExecuteCommand("apt-get remove --purge nodejs -y")
		utils.ExecuteCommand("rm " + gpgFile + " /etc/apt/sources.list.d/nodesource.list")
	}
	version := selectVersion()

	// extra packages
	var packages []string
	if !utils.CheckPackage("apt-transport-https") {
		packages = append(packages, "apt-transport-https")
	}
	if !utils.CheckPackage("lsb-release") {
		packages = append(packages, "lsb-release")
	}
	if !utils.CheckPackage("gpg") {
		packages = append(packages, "gnupg")
	}
	if !utils.CheckPackage("build-essential") {
		packages = append(packages, "build-essential")
	}
	if len(packages) > 0 {
		utils.ExecuteCommand("apt-get install " + strings.Join(packages, " ") + " -y")
	}

	utils.ExecuteCommand("curl -s https://deb.nodesource.com/gpgkey/nodesource.gpg.key | gpg --dearmor | tee " + gpgFile + " >/dev/null")
	utils.ExecuteCommand("echo 'deb [signed-by=" + gpgFile + "] https://deb.nodesource.com/node_" + version + ".x bullseye main' > /etc/apt/sources.list.d/nodesource.list")
	utils.ExecuteCommand("echo 'deb-src [signed-by=" + gpgFile + "] https://deb.nodesource.com/node_" + version + ".x bullseye main' >> /etc/apt/sources.list.d/nodesource.list")
	utils.ExecuteCommand("apt-get update -y")
	utils.ExecuteCommand("apt-get install nodejs -y")

	prompt := promptui.Select{
		Label:    "Voulez-vous installer yarn (alternative à npm)",
		Items:    []string{"Oui", "Non"},
		HideHelp: true,
	}
	_, result, err := prompt.Run()
	if err != nil {
		utils.ShowMessage("Une erreur est survenue !", true)
	}
	if result == "Oui" {
		yarnInstall()
	}
	utils.GoodBye()
}
