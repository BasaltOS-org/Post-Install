package packages

import (
	"PostInstall/utils"
	"fmt"
	"os/exec"
)

type Packages []string

func Install(p Packages) error {
	// assume running as root
	utils.Logger.Info("Installing Packages", "Packages", p) 	
	arguments := []string{"install", "-y"}
	arguments = append(arguments, p...)
	fmt.Println(arguments)

	cmd := exec.Command("/usr/bin/dnf", arguments...)
	

	_, err := cmd.CombinedOutput()
	if err != nil {
		utils.Logger.Error("error returned", "error", err)
		return err
	}


	utils.Logger.Info("Installed Packages", "Packages", p)	
	return nil
}


func Remove(p Packages) error {
	// assume user is running as root
		arguments := []string{"remove", "-y"}
		arguments = append(arguments, p...)
		fmt.Println(arguments)

		cmd := exec.Command("/usr/bin/dnf", arguments...)

		if _, err := cmd.CombinedOutput(); err != nil {
			utils.Logger.Error("Error returned", "error", err)
			return err
		}
	utils.Logger.Info("Removed Package Group with", "Packages", p)
	return nil
}

