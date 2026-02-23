package cmd

import (
	"os"

	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/util"

	"github.com/spf13/cobra"
)

var signerCmd = &cobra.Command{
	Use:                "signer",
	Short:              "Passes args through to signer app.",
	Long:               `Passes args through to signer app.`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, _ []string) {
		args := util.GetCommandArgs(cmd)
		if len(args) > 0 && args[0] == "-" {
			args[0] = "signer"
		}
		exitCode, _ := apps.Signer.Execute(args...)
		os.Exit(exitCode)
	},
}

func init() {
	signerCmd.Flags().SetInterspersed(false)

	RootCmd.AddCommand(signerCmd)
}
