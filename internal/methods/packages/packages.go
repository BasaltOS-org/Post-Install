package packages

import (
	"PostInstall/internal/database"
	"PostInstall/utils"
	"database/sql"
	"fmt"
	"os/exec"
)



type PackageGrouper interface{
	View() string
	Install() (bool, error)
	Delete() (bool, error)
}

type PackageGroup struct {
	PackageID int
	Name string
	Packages []string
}



func (p PackageGroup) Install() error {
	// These should always be run as sudo, It is the job of the frontend of choice to provide those privelages
	utils.Logger.Info("Executed Install", "pkgName", p.Name, "pkgID", p.PackageID)

	if value, err := database.QueryPackageStatus(p.PackageID); value == true {
		fmt.Println("Package Group Already installed!")
		utils.Logger.Warn("Package Group Already Installed, Aborted Install", "Group", p.Name)
		return err
	}

	utils.Logger.Info("Installing Package Group", "Group", p.Name) // TODO: Make this pretty text so it stands out
	for _, pkg := range p.Packages { // Install each package one by one since that's less error prone
		cmd := exec.Command("sudo", "dnf", "install", "-y", pkg) // -y assumes yes and doesn't prompt for confirm
		fmt.Println("Installing Package:", pkg)

		_, err := cmd.CombinedOutput()
		if err != nil {
			utils.Logger.Error("error returned", "error", err)
			return err
		}

		fmt.Println("Installed Package", pkg)
	}


	if err := database.InsertPackageStatus(p.PackageID, true); err != nil {
		utils.Logger.Error("error returned", "error", err)
		return err
	}
	utils.Logger.Info("Installed Package Group", "Group Name", p.Name, "Packages", p.Packages)	
	return nil
}


func (p PackageGroup) View() PackageGroup {
	utils.Logger.Info("Executed View() for","Group", p.Name)
	return p
}

func (p PackageGroup) Remove() error {
	// assume user is running as root
	utils.Logger.Info("Executed Delete()", "Group", p.Name)
	if value, err := database.QueryPackageStatus(p.PackageID); err == sql.ErrNoRows || value == false {
		fmt.Println("Package Group Not installed!")
		utils.Logger.Warn("Package Group Not Installed, Aborted Remove", "Group", p.Name)
		return err
	}

	for _, pkg := range p.Packages {	
		cmd := exec.Command("sudo", "dnf", "remove", "-y",pkg )

		utils.Logger.Info("Removing Package", "package", pkg)
		fmt.Println("Removing Package:", pkg)

		_, err := cmd.CombinedOutput() 
		
		if err != nil {
			utils.Logger.Error("Error returned", "error", err)
			return err
		}
		
		err = database.UpdatePackageStatus(p.PackageID, false) 
		if err != nil {
			utils.Logger.Error("Error returned while updating table values", "error", err)
			return err
		}
		fmt.Println("Removed Package", pkg)
		utils.Logger.Info("Removed Package", "package", pkg)

	}
	utils.Logger.Info("Deleted Package Group", "Group Name", p.Name, "Packages", p.Packages)
	return nil
}

