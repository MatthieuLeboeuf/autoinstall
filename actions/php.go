package actions

import (
	"github.com/MatthieuLeboeuf/autoinstall/utils"
	"github.com/gocolly/colly/v2"
	"github.com/manifoldco/promptui"
	"strings"
)

func selectPhpVersion() string {
	var phpVersions []string

	c := colly.NewCollector()
	c.OnHTML("table tbody", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, row *colly.HTMLElement) {
			if strings.Contains(row.Attr("href"), "/downloads.php#") {
				phpVersions = append(phpVersions, row.Text)
			}
		})
	})
	err := c.Visit("https://www.php.net/supported-versions")
	if err != nil {
		utils.ShowMessage("Une erreur est survenue !", true)
	}

	prompt := promptui.Select{
		Label:    "Quelle version voulez-vous",
		Items:    phpVersions,
		HideHelp: true,
	}
	_, result, err := prompt.Run()
	if err != nil {
		utils.ShowMessage("Une erreur est survenue !", true)
	}
	return result
}

func changeDefaultPhpVersion() bool {
	prompt := promptui.Select{
		Label:    "Voulez-vous mettre cette version par défaut",
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

func installRecommendedPhpExtensions() bool {
	prompt := promptui.Select{
		Label:    "Voulez-vous installer les extensions php fréquemment installés",
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

func PhpInstall() {
	version := selectPhpVersion()

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
	if !utils.CheckPackage("ca-certificates") {
		packages = append(packages, "ca-certificates")
	}
	if len(packages) > 0 {
		utils.ExecuteCommand("apt-get install " + strings.Join(packages, " ") + " -y")
	}

	gpgFile := "/usr/share/keyrings/deb.sury.org-php.gpg"
	utils.ExecuteCommand("curl -s https://packages.sury.org/php/apt.gpg | gpg --dearmor | tee " + gpgFile + " >/dev/null")
	utils.ExecuteCommand("echo 'deb [signed-by=" + gpgFile + "] https://packages.sury.org/php/ " + utils.GetOsName() + " main' > /etc/apt/sources.list.d/php.list")
	utils.ExecuteCommand("apt-get update -y")
	utils.ExecuteCommand("apt-get install php" + version + "-fpm -y")

	if changeDefaultPhpVersion() {
		utils.ExecuteCommand("update-alternatives --set php /usr/bin/php" + version)
	}

	if installRecommendedPhpExtensions() {
		utils.ExecuteCommand("apt-get install php" + version + "-{curl,mbstring,gd,xml,zip,intl,mysql} -y")
	}

	utils.GoodBye()
}
