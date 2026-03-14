package cmd

import (
	"BasaltPostInstallAssistant/internal/methods/packages"
	"BasaltPostInstallAssistant/utils"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func cmdPackage() *cli.Command {
	cmd := &cli.Command{
		Name: "package",
		Aliases: []string{"pkg"},
		Description: "Allows you to work with groups of packages that have been handpicked by the developers for convinience",
		UsageText: "bia package [subcommmand] [flags]",
		Commands: []*cli.Command{
			packageInstall(),
			packageRemove(),
		},
	}
	return cmd
}

func packageInstall() *cli.Command {
	cmd := &cli.Command{
		Name: "install",
		Aliases: []string{"in", "i"},
		Usage: "Use this command to install groups of packages",
		Description: "This command is used to install groups of packages, Run the Command TBD to view the available list of packages\n" +
		"all packages installed are from the official repositories and are handpicked by the developers to ensure compatibility.",
		Action: func(ctx context.Context, c *cli.Command) error {
			pkgGrp := os.Args[3] // this should only be the PackageGroup they want to install

			pkg, err := packages.DeterminePkg(pkgGrp)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			err = pkg.Install()
			if err != nil {
				fmt.Println("error encountered", err)
				utils.Logger.Error("error encountered", "error", err)
				os.Exit(1)
			}
			utils.Logger.Info("successfully installed package")
			return nil
		},
		
		

	}
	return cmd
}

func packageRemove() *cli.Command {
	cmd := &cli.Command{
		Name: "remove",
		Aliases: []string{"rm"},
		Description: "This command is used to uninstall groups of packages, Run the Command TBD to view installed package groups",
		UsageText: "package remove [packageGroup]... [flags]",
		Action: func(ctx context.Context, c *cli.Command) error {
			fmt.Println("removed")
			return  nil
		},
	}
	return cmd
}