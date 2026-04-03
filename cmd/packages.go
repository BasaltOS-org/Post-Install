package cmd

import (
	"BasaltPostInstallAssistant/internal/methods/packages"
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

			for _, i := range pkgGrp {
				pkg, err := packages.DeterminePkg(i)
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}

				if err = pkg.Install(); err != nil {
					fmt.Println(err)
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
			
			for _, i := range pkgGrp {
				pkg, err := packages.DeterminePkg(i)
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
				if err = pkg.Remove(); err != nil {
					fmt.Println(err)
				}
			}
		return  nil
		},
	}

	return cmd
}

func packageList() *cli.Command {
	cmd := &cli.Command{
		Name: "list",
		Aliases: []string{"l", "show", "sh"},
		Flags: []cli.Flag{
			
			&cli.BoolFlag {
				Name: "installed",
				Value: true,
				Aliases: []string{"in"},
			},
		},
		Description: "This command is used to View package groups, Use flags to filter",
		UsageText: "package remove [packageGroup]... [flags]",
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.Bool("installed") == true {
			} else {
				
			}
			return nil
		},
	}
	return cmd
}
