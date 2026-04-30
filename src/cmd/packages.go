package cmd

import (
	"PostInstall/internal/database"
	"PostInstall/internal/methods/packages"
	"PostInstall/utils"
	"context"
	"fmt"


	"github.com/urfave/cli/v3"
)

func cmdPackage() *cli.Command {
	cmd := &cli.Command{
		Name:        "package",
		Aliases:     []string{"pkg"},
		Description: "Allows you to work with groups of packages that have been handpicked by the developers for convinience",
		UsageText:   "bia package [subcommmand] [flags]",
		Commands: []*cli.Command{
			packageInstall(),
			packageRemove(),
			packageList(),
		},
	}
	return cmd
}

func packageInstall() *cli.Command {
	cmd := &cli.Command{
		Name:    "install",
		Aliases: []string{"in", "i"},
		Usage:   "Use this command to install groups of packages",
		Description: "This command is used to install groups of packages, Run the Command packages list to view the available list of packages\n" +
			"all packages installed are from the official repositories and are handpicked by the developers to ensure compatibility.",
		UsageText: "package install [packageGroup]... [flags]",
		Action: func(ctx context.Context, c *cli.Command) error {
			pkgGroups := c.Args().Slice()
			if len(pkgGroups) <= 0 {
				utils.Logger.Error("packageInstall didn't recieve enough arguments")
			}

			for _, val := range pkgGroups {
				pkg, err := database.GetPackageGroup(val)
				if err == database.ErrNotFound {
					utils.Logger.Error("packageInstall: ", "error", err)
					fmt.Println(err)
					return err
				}

				if pkg.Installed == true {
					utils.Logger.Warn("Package Already Installed, Aborting Installation Now")
					fmt.Println("Package Already Installed, Aborting Installation")
					return nil
				}

				if err := packages.Install(&pkg); err != nil {
					utils.Logger.Error("packageInstall: ", "error", err)
					fmt.Println(err)
					return err
				}
				if err := database.UpdateInstalledStatus(val, true); err != nil {
					utils.Logger.Error("packageInstall: ", "error", err)
					fmt.Println(err)
					return err
				}
			}
			return nil
		},
	}
	return cmd
}

func packageRemove() *cli.Command {
	cmd := &cli.Command{
		Name:        "remove",
		Aliases:     []string{"rm"},
		Description: "This command is used to uninstall groups of packages, Run the Command TBD to view installed package groups",
		UsageText:   "package remove [packageGroup]... [flags]",
		Action: func(ctx context.Context, c *cli.Command) error {
			pkgGroups := c.Args().Slice()
			if len(pkgGroups) <= 0 {
				utils.Logger.Error("packageRemove didn't recieve enough arguments")
			}


			for _, val := range pkgGroups {
				pkg, err := database.GetPackageGroup(val) 
				if err == database.ErrNotFound {
					utils.Logger.Error("packageInstall: ", "error", err)
					fmt.Println(err)
					return err
				}

				if pkg.Installed != true {
					utils.Logger.Warn("Package is not Installed, Aborting Removal Now")
					fmt.Println("Package is not installed, Aborting Removal")
					return nil
				}

				if err := packages.Remove(&pkg); err != nil {
					utils.Logger.Error("packageInstall: ", "error", err)
					fmt.Println(err)
					return err
				}
				err = database.UpdateInstalledStatus(val, false)
				if err != nil {
					utils.Logger.Error("packageInstall: ", "error", err)
					fmt.Println(err)
					return err
				}
			}
			return nil
		},
	}

	return cmd
}

func packageList() *cli.Command {
	cmd := &cli.Command{
		Name:        "list",
		Aliases:     []string{"ls", "show", "sh"},
		Description: "This command is used to View package groups that can be installed",
		UsageText:   "package list [packageGroup]... [flags]",
		Action: func(ctx context.Context, c *cli.Command) error {

			pmap := database.ListPackages()

			for key, val := range pmap {
				fmt.Printf("Package Group Name: %v holds Packages: %+v is Installed: %v\n", key, val.Packages, val.Installed)
			}
			return nil
		},
	}
	return cmd
}
