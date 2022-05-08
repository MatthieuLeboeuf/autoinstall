package actions

import (
	"github.com/MatthieuLeboeuf/autoinstall/utils"
	"github.com/manifoldco/promptui"
	"github.com/sethvargo/go-password/password"
)

func exposeMariadbAllInterfaces() bool {
	prompt := promptui.Select{
		Label:    "Voulez-vous exposer MariaDB sur toutes les interfaces",
		Items:    []string{"Oui", "Non"},
		HideHelp: true,
	}
	_, result, err := prompt.Run()
	if err != nil {
		utils.ShowMessage("Une erreur est survenue !", true)
	}
	if result == "Non" {
		return false
	}
	return true
}

func MariaDbInstall() {
	if utils.CheckPackage("mariadb-common") {
		utils.ShowMessage("MariaDB est déja installé sur le système", false)
		prompt := promptui.Select{
			Label:    "Voulez-vous le réinstaller",
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
		utils.ExecuteCommand("apt-get remove --purge mariadb-server mariadb-client mariadb-common -y")
	}

	utils.ExecuteCommand("apt-get install mariadb-server mariadb-client mariadb-common -y")

	rootPassword, err := password.Generate(50, 10, 0, false, false)
	if err != nil {
		utils.ShowMessage("Une erreur est survenue !", true)
	}

	utils.ExecuteCommand("mysql -e \"UPDATE mysql.user SET Password = PASSWORD('" + rootPassword + "') WHERE User = 'root'\"")
	utils.ExecuteCommand("mysql -e \"DROP USER ''@'localhost'\"")
	utils.ExecuteCommand("mysql -e \"DROP USER ''@'" + utils.GetHostname() + "'\"")
	utils.ExecuteCommand("mysql -e \"DROP DATABASE test\"")
	utils.ExecuteCommand("mysql -e \"FLUSH PRIVILEGES\"")

	if exposeMariadbAllInterfaces() {
		utils.ExecuteCommand("sed -i 's/127.0.0.1/0.0.0.0/g' /etc/mysql/mariadb.conf.d/50-server.cnf")
		utils.ExecuteCommand("service mariadb restart")
	}

	utils.ShowMessage("L'installation de MariaDB à été effectuée", false)
	utils.ShowMessage("Mot de passe root : "+rootPassword, false)

	utils.GoodBye()
}
