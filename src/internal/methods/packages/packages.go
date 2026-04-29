package packages

import (
	"PostInstall/utils"
	"fmt"
	"os/exec"
)



type PackageGroup struct {
	PackageID int
	Name string
	Packages []string
}



func (p PackageGroup) Install() error {

	// These should always be run as sudo, It is the job of the frontend of choice to provide those privelages
	utils.Logger.Info("Executed Install", "pkgName", p.Name, "pkgID", p.PackageID)


	utils.Logger.Info("Installing Packages", "Packages", p.Packages) 
	for _, pkg := range p.Packages { // Install each package one by one since that's less error prone
		cmd := exec.Command("dnf", "install", "-y", pkg) // -y assumes yes and doesn't prompt for confirm
		fmt.Println("Installing Package:", pkg)

		_, err := cmd.CombinedOutput()
		if err != nil {
			utils.Logger.Error("error returned", "error", err)
			return err
		}

		fmt.Println("Installed Package", pkg)
	}

	utils.Logger.Info("Installed Package Group", "Group Name", p.Name, "Packages", p.Packages)	
	return nil
}


func (p PackageGroup) Remove() error {
	// assume user is running as root
	utils.Logger.Info("Executed Delete()", "Group", p.Name)
	
	for _, pkg := range p.Packages {	
		cmd := exec.Command("dnf", "remove", "-y",pkg )

		utils.Logger.Info("Removing Package", "package", pkg)
		fmt.Println("Removing Package:", pkg)

		_, err := cmd.CombinedOutput() 
		
		if err != nil {
			utils.Logger.Error("Error returned", "error", err)
			return err
		}
		

		fmt.Println("Removed Package", pkg)
		utils.Logger.Info("Removed Package", "package", pkg)

	}
	utils.Logger.Info("Deleted Package Group", "Group Name", p.Name, "Packages", p.Packages)
	return nil
}

