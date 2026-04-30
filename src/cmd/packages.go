package cmd

import (
	"PostInstall/internal/database"
	"PostInstall/internal/methods/packages"
	"context"
	"fmt"

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
			packageList(),
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
			pkgGrp := c.Args().Slice()

			for _, val := range pkgGrp {
				pkg := database.GetPackage(val)

				if err := packages.Install(pkg); err != nil {
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
		Name: "remove",
		Aliases: []string{"rm"},
		Description: "This command is used to uninstall groups of packages, Run the Command TBD to view installed package groups",
		UsageText: "package remove [packageGroup]... [flags]",
		Action: func(ctx context.Context, c *cli.Command) error {
			pkgGrp := c.Args().Slice()
			
			for _, val := range pkgGrp {
				pkg := database.GetPackage(val)

				if err := packages.Remove(pkg); err != nil {
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
		Name: "list",
		Aliases: []string{"ls", "show", "sh"},
		Description: "This command is used to View package groups that can be installed",
		UsageText: "package remove [packageGroup]... [flags]",
		Action: func(ctx context.Context, c *cli.Command) error {

			pmap := database.ListPackages()

			for key, val := range pmap {
				fmt.Printf("Package Group Name: %v For Packages: %+v\n", key, val)
			}
			return nil
		},
	}
	return cmd
}
