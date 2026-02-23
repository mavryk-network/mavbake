package cmd

import (
	"os"
	"slices"

	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/util"

	"github.com/spf13/cobra"
)

var mavkitSignerCmd = &cobra.Command{
	Use:                "mavkit-signer",
	Short:              "Passes args through to signer app - mavkit-signer.",
	Long:               `Passes args through to signer app - mavkit-signer.`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, _ []string) {
		args := util.GetCommandArgs(cmd)
		args = slices.Insert(args, 0, "signer")
		exitCode, _ := apps.Signer.Execute(args...)
		os.Exit(exitCode)
	},
}

func init() {
	mavkitSignerCmd.Flags().SetInterspersed(false)

	RootCmd.AddCommand(mavkitSignerCmd)
}
