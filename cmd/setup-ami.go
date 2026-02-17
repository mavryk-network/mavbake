package cmd

import (
	"os"

	"github.com/mavryk-network/mavbake/ami"
	"github.com/mavryk-network/mavbake/system"
	"github.com/mavryk-network/mavbake/util"
	"go.alis.is/common/log"

	"github.com/spf13/cobra"
)

var setupAmiCmd = &cobra.Command{
	Use:   "setup-ami",
	Short: "Install ami and eli.",
	Long:  "Install latest ami and eli.",
	Run: func(cmd *cobra.Command, args []string) {
		system.RequireElevatedUser()

		exitCode, err := ami.Install(util.GetCommandBoolFlag(cmd, "silent"))
		if err != nil {
			log.Error("Failed to install ami and eli!", "error", err)
			os.Exit(exitCode)
		}
	},
}

func init() {
	setupAmiCmd.Flags().Bool("silent", false, "Do not print any output.")
	RootCmd.AddCommand(setupAmiCmd)
}
