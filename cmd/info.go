package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/cli"
	"github.com/mavryk-network/mavbake/constants"
	"github.com/mavryk-network/mavbake/util"
	"go.alis.is/common/log"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Prints runtime information about MavBake.",
	Long:  "Collects and prints runtime information about MavBake instance.",
	Run: func(cmd *cobra.Command, args []string) {
		timeout, _ := cmd.Flags().GetInt("timeout")
		if timeout <= 0 {
			timeout = 5
		}

		result := map[string]any{}

		for _, v := range GetAppsBySelectionCriteria(cmd, AppSelectionCriteria{
			InitialSelection:  InstalledApps,
			FallbackSelection: ImplicitApps,
			OptionCheckType:   InfoOptionCheck,
		}) {
			options := map[string]any{
				"timeout": timeout,
			}
			for _, option := range v.GetAvailableInfoCollectionOptions() {
				switch option.Type {
				case "bool":
					if checked, _ := cmd.Flags().GetBool(fmt.Sprintf("%s-%s", v.GetId(), option.Name)); checked {
						options[option.Name] = true
					}
				}
			}

			log.Debug("Collecting info for:", "app", v.GetId())
			optionsJson, _ := json.Marshal(options)
			if cli.JsonLogFormat {
				result[v.GetId()], _ = v.GetInfo(optionsJson)
			} else {
				err := v.PrintInfo(optionsJson)
				util.AssertE(err, fmt.Sprintf("Failed to collect %s's info!", v.GetId()))
			}
		}

		if cli.JsonLogFormat {
			output, err := json.Marshal(result)
			util.AssertEE(err, "Failed to serialize MavBake runtime info!", constants.ExitSerializationFailed)
			fmt.Println(string(output))
			return
		}
	},
}

func init() {
	infoCmd.Flags().Int("timeout", 5, "How long to wait for collecting info.")
	for _, v := range apps.All {
		infoCmd.Flags().Bool(v.GetId(), false, fmt.Sprintf("Prints info for %s.", v.GetId()))
		for _, option := range v.GetAvailableInfoCollectionOptions() {
			switch option.Type {
			case "bool":
				infoCmd.Flags().Bool(fmt.Sprintf("%s-%s", v.GetId(), option.Name), false, "")
			}
		}
	}
	RootCmd.AddCommand(infoCmd)
}
