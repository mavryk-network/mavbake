package cmd

import (
	"os"

	"github.com/mavryk-network/mavbake/apps"

	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Passes args through to node app.",
	Long:  `Passes args through to node app.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && args[0] == "-" {
			args[0] = "node"
		}
		exitCode, _ := apps.Node.Execute(args...)
		os.Exit(exitCode)
	},
}

func init() {
	nodeCmd.Flags().SetInterspersed(false)

	RootCmd.AddCommand(nodeCmd)
}
