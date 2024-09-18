package cmd

import (
	"fmt"

	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/system"
	"github.com/mavryk-network/mavbake/util"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops BB.",
	Long:  "Stops services of BB instance.",
	Run: func(cmd *cobra.Command, args []string) {
		system.RequireElevatedUser()

		for _, v := range GetAppsBySelectionCriteria(cmd, AppSelectionCriteria{
			InitialSelection:  InstalledApps,
			FallbackSelection: ImplicitApps,
		}) {
			exitCode, err := v.Stop()
			util.AssertEE(err, fmt.Sprintf("Failed to stop %s's services!", v.GetId()), exitCode)
		}

		log.Info("Requested services stopped succesfully")
	},
}

func init() {
	for _, v := range apps.All {
		stopCmd.Flags().Bool(v.GetId(), false, fmt.Sprintf("Stops %s's services.", v.GetId()))
	}
	RootCmd.AddCommand(stopCmd)
}
