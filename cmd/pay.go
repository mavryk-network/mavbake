package cmd

import (
	"os"
	"strings"

	"github.com/samber/lo"
	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/constants"
	"github.com/mavryk-network/mavbake/util"

	"github.com/spf13/cobra"
)

var payCmd = &cobra.Command{
	Use:                "pay",
	Short:              "Passes args through to mavpay app.",
	Long:               `Passes args through to mavpay app.`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, _ []string) {
		args := util.GetCommandArgs(cmd)
		util.AssertBE(apps.Pay.IsInstalled(), "Pay app is not installed!", constants.ExitAppNotInstalled)
		nonOptionArgsCount := lo.CountBy(args, func(s string) bool { return !strings.HasPrefix(s, "-") })
		hasHelpOption := lo.ContainsBy(args, func(s string) bool { return s == "-h" || s == "--help" })

		if nonOptionArgsCount == 0 && !hasHelpOption {
			args = append([]string{"pay"}, args...) // default to "pay" subcommand
		}
		if len(args) > 0 && args[0] == "-" {
			args[0] = "pay"
		}
		exitCode, _ := apps.Pay.Execute(args...)
		os.Exit(exitCode)
	},
}

func init() {
	payCmd.Flags().SetInterspersed(false)

	RootCmd.AddCommand(payCmd)
}
