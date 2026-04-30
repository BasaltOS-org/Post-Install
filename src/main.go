package main

import (
	"PostInstall/cmd"
	"PostInstall/utils"
	"context"
	"fmt"
	"os"

	sys "golang.org/x/sys/unix"

	"github.com/urfave/cli/v3"
)

var ErrNotRoot = fmt.Errorf("User is not root")

func init() {
	utils.InitLogger()
}

func privCheck() bool {
	if id := sys.Getuid(); id != 0 {
		return false
	}
	return true
}

func main() {
	// Initialize and show the UI
	//ui.ShowUI()

	// Temporarily, Let's Forget the UI exists just so i can focus on the CLI

	isRoot := privCheck()
	if isRoot != true {
		utils.Logger.Error("error encounted", "error", ErrNotRoot)
		fmt.Println(ErrNotRoot)
		os.Exit(1)
	}

	cmd := &cli.Command{
		Commands: cmd.InitCmds(),
	}
	cmd.Run(context.Background(), os.Args)

}
