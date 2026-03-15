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



func (p PackageGroup) Install() (err error){
	// These should always be run as sudo, It is the job of the frontend of choice to provide those privelages
	utils.Logger.Info(fmt.Sprintf("Executed Install for %v", p.Name))

	for _, pkg := range p.Packages { // Install each package one by one since that's less error prone
		cmd := exec.Command("sudo", "dnf", "install", "-y", pkg) // -y assumes yes and doesnt prompt for confirm
		utils.Logger.Info("Installing Package Group", "Group", p.Name)

		// cmd.CombinedOutput() returns the output and error (if any), in the future output itself would be useless (in the context of the gui)
		// but for ease of development it will be logged but not returned, in the future this value could be discarded
		if out, err := cmd.CombinedOutput(); err != nil { 
			utils.Logger.Error("error returned", "error", err)
			utils.Logger.Error(fmt.Sprintf("output for debug %v \t", string(out)))
			return err
		}
		utils.Logger.Info("Installed Package Group", "Group Name", p.Name, "Packages", p.Packages)
			
	}
	return nil
}


func (p PackageGroup) View() PackageGroup {
	utils.Logger.Info("Executed View() for","Group", p.Name)
	return p
}

func (p PackageGroup) Remove() (err error){
	// assume user is running as root
	utils.Logger.Info("Executed Delete()", "Group", p.Name)
	
	for _, pkg := range p.Packages {	
		cmd := exec.Command("sudo", "dnf", "remove", "-y",pkg )
		utils.Logger.Info("Removing Package", "package", pkg)
		fmt.Println("Removing Package: ", pkg)

		if out, err := cmd.CombinedOutput(); err != nil {
			utils.Logger.Error("Error returned", "error", err)
			utils.Logger.Error(fmt.Sprintf("Output for debug %v", string(out)))
			return err
		}
		utils.Logger.Info("Deleted Package Group", "Group Name", p.Name, "Packages", p.Packages)

	}
	return nil
}

