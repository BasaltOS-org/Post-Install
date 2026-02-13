package main

import (
	"BasaltPostInstallAssistant/internal/methods/packages"
	"BasaltPostInstallAssistant/utils"
)

func init() {
	utils.InitLogger()
}

func main() {
	var TestPackage packages.PackageGroup = packages.PackageGroup{
		Name:     "test",
		Packages: []string{"vim", "go", "lua"},
	}
	TestPackage.Install()
}
