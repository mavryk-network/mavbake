package cmd

import (
	"os"

	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/util"

	"github.com/spf13/cobra"
)

var dalCmd = &cobra.Command{
	Use:                "dal",
	Short:              "Passes args through to dal node app.",
	Long:               `Passes args through to dal node app.`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, _ []string) {
		args := util.GetCommandArgs(cmd)
		if len(args) > 0 && args[0] == "-" {
			args[0] = "dal-node"
		}
		exitCode, _ := apps.DalNode.Execute(args...)
		os.Exit(exitCode)
	},
}

func init() {
	dalCmd.Flags().SetInterspersed(false)
	dalCmd.Hidden = true

	RootCmd.AddCommand(dalCmd)
}
