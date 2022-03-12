package main

import (
	"fmt"
	"github.com/MatthieuLeboeuf/autoinstall/actions"
	"github.com/MatthieuLeboeuf/autoinstall/utils"
	. "github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"os/user"
)

func main() {
	fmt.Println(Cyan("Développé par Matthieu Leboeuf").Italic())
	fmt.Println(Cyan("AutoUpdater").Bold(), "- Lancement")
	CheckPermissions()
	CheckRequirements()
	SelectAction()
}

func CheckPermissions() {
	currentUser, err := user.Current()
	if err != nil {
		utils.ShowMessage("Erreur !", true)
	}
	if currentUser.Username != "root" {
		utils.ShowMessage("Erreur, Vous n'avez pas les permissions nécessaires !", true)
	}
}

func CheckRequirements() {
	OSInfo := utils.ReadOSRelease()
	if OSInfo["ID"] != "debian" || OSInfo["VERSION_ID"] != "11" {
		utils.ShowMessage("Cette app est uniquement compatible avec debian 11 !", true)
	}
	prompt := promptui.Select{
		Label:    "Voulez-vous faire les mises à jour du système",
		Items:    []string{"Oui", "Non"},
		HideHelp: true,
	}
	_, result, err := prompt.Run()
	if err != nil {
		utils.ShowMessage("Une erreur est survenue !", true)
	}
	utils.ShowMessage("Execution de apt update", false)
	utils.ExecuteCommand("apt-get update -y")
	utils.ShowMessage("Execution de apt autoremove", false)
	utils.ExecuteCommand("apt-get autoremove -y")
	if !utils.CheckPackage("curl") {
		utils.ShowMessage("Execution de apt install curl", false)
		utils.ExecuteCommand("apt-get install curl -y")
	}
	if result == "Oui" {
		utils.ShowMessage("Execution de apt full-upgrade", false)
		utils.ExecuteCommand("apt-get full-upgrade -y")
	}
}

func SelectAction() {
	prompt := promptui.Select{
		Label:    "Sélectionner ce que vous voulez faire",
		Items:    []string{"Installer NodeJs"},
		HideHelp: true,
	}
	_, result, err := prompt.Run()
	if err != nil {
		utils.ShowMessage("Une erreur est survenue !", true)
	}
	switch result {
	case "Installer NodeJs":
		actions.NodeJsInstall()
	}
}
