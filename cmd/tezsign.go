package cmd

import (
	"os"
	"slices"

	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/util"

	"github.com/spf13/cobra"
)

var mavsignCmd = &cobra.Command{
	Use:                "mavsign",
	Short:              "Passes args through to signer app - mavsign.",
	Long:               `Passes args through to signer app - mavsign.`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, _ []string) {
		args := util.GetCommandArgs(cmd)
		args = slices.Insert(args, 0, "mavsign")
		exitCode, _ := apps.Signer.Execute(args...)
		os.Exit(exitCode)
	},
}

func init() {
	mavsignCmd.Flags().SetInterspersed(false)

	RootCmd.AddCommand(mavsignCmd)
}
