package cmd

import (
	"fmt"

	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/system"
	"github.com/mavryk-network/mavbake/util"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts BB.",
	Long:  "Starts services of BB instance.",
	Run: func(cmd *cobra.Command, args []string) {
		system.RequireElevatedUser()

		for _, v := range GetAppsBySelectionCriteria(cmd, AppSelectionCriteria{
			InitialSelection:  InstalledApps,
			FallbackSelection: ImplicitApps,
		}) {
			exitCode, err := v.Start()
			util.AssertEE(err, fmt.Sprintf("Failed to starts %s's services!", v.GetId()), exitCode)
		}

		log.Info("Requested services started succesfully")
	},
}

func init() {
	for _, v := range apps.All {
		startCmd.Flags().Bool(v.GetId(), false, fmt.Sprintf("Starts %s's services.", v.GetId()))
	}
	RootCmd.AddCommand(startCmd)
}
