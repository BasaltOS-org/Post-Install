package packages

import (
	"BasaltPostInstallAssistant/utils"
	"fmt"
	"os/exec"
)



type PackageGrouper interface{
	View() string
	Install() (bool, error)
	Delete() (bool, error)
}

type PackageGroup struct {
	Name string
	Packages []string
}



func (p PackageGroup) Install() (success bool, err error){
	// TODO: Use Polkit or the cli frontend depending on the context to check / ask for sudo perms when applicable
	// For Now lets assume the user is always running as root
	utils.Logger.Info(fmt.Sprintf("Executed Install for %v", p.Name))

	for _, pkg := range p.Packages { // Install each package one by one since that's less error prone
		cmd := exec.Command("dnf", "install", "-y", pkg) // -y assumes yes and doesnt prompt for confirm
		utils.Logger.Info(fmt.Sprintf("Installing Package from Group: %v", p.Name))

		// cmd.CombinedOutput() returns the output and error (if any), in the future output itself would be useless (in the context of the gui)
		// but for ease of development it will be logged but not returned, in the future this value could be discarded
		if out, err := cmd.CombinedOutput(); err != nil { 
			utils.Logger.Error(fmt.Sprintf("error returned %v", err))
			utils.Logger.Error(fmt.Sprintf("output for debug %v \t", string(out)))
			return false, err
		}
			
	}
	return true, err
}


func (p PackageGroup) View() PackageGroup {
	utils.Logger.Info(fmt.Sprintf("Executed View() for: %v", p.Name))
	return p
}

func (p PackageGroup) Delete() (success bool, err error){
// TODO: Use Polkit or the cli frontend depending on the context to check / ask for sudo perms when applicable
// For Now lets assume the user is always running as root
	utils.Logger.Info(fmt.Sprintf("Executed Delete() for %v", p.Name))
	utils.Logger.Info(fmt.Sprintf(" for %v", p.Name))
	
	for _, pkg := range p.Packages {	
		cmd := exec.Command("dnf", "remove", "-y",pkg )

		if out, err := cmd.CombinedOutput(); err != nil {
			utils.Logger.Error(fmt.Sprintf("Error returned %v", err))
			utils.Logger.Error(fmt.Sprintf("Output for debug %v", string(out)))
			return false, err
		}
		utils.Logger.Info(fmt.Sprintf("Deleting Packages from Group: %v", p.Name))
		utils.Logger.Info(fmt.Sprintf("Deleting Packages: %+v", p.Packages))

	}
	return true, nil
}

