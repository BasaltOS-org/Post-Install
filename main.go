package main

import (
	"BasaltPostInstallAssistant/internal/methods/packages"
	"BasaltPostInstallAssistant/ui"
	"BasaltPostInstallAssistant/utils"
)

func init() {
	utils.InitLogger()
}

func main() {
	// Initialize and show the UI
	ui.ShowUI()

	// The following code will run after the UI is closed
	var TestPackage packages.PackageGroup = packages.PackageGroup{
		Name:     "test",
		Packages: []string{"vim", "go", "lua"},
	}
	TestPackage.Install()
}
