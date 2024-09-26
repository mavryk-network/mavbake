package cmd

import (
	"fmt"

	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/apps/base"
	"github.com/samber/lo"

	"github.com/spf13/cobra"
)

var isAppInstalledCmd = &cobra.Command{
	Use:    "is-app-installed",
	Short:  "Checks whether specific app is instaleled BB.",
	Long:   "Stops services of BB instance.",
	Hidden: true,
	Args:   cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		app, found := lo.Find(apps.All, func(v base.MavPayApp) bool {
			return v.GetId() == id
		})
		fmt.Println(found && app.IsInstalled())
	},
}

func init() {
	RootCmd.AddCommand(isAppInstalledCmd)
}
