package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/cli"
	"github.com/mavryk-network/mavbake/constants"
	"github.com/mavryk-network/mavbake/util"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Prints BB CLI apps.",
	Long:  "Prints MavPay CLI apps.",
	Run: func(cmd *cobra.Command, args []string) {
		appsTable := table.NewWriter()
		appsTable.SetOutputMirror(os.Stdout)
		appsTable.SetStyle(table.StyleLight)
		appsTable.AppendHeader(table.Row{"App", "Installed?"}, table.RowConfig{AutoMerge: true})

		result := map[string]interface{}{}
		for _, v := range apps.All {
			isInstalled := v.IsInstalled()
			result[v.GetId()] = map[string]interface{}{
				"installed": isInstalled,
			}
			appsTable.AppendRow(table.Row{v.GetId(), isInstalled})
		}

		if cli.JsonLogFormat || cli.IsRemoteInstance {
			data, err := json.Marshal(result)
			util.AssertEE(err, "Failed to serialize apps info!", constants.ExitSerializationFailed)
			fmt.Println(string(data))
			return
		}

		appsTable.Render()
	},
}

func init() {
	RootCmd.AddCommand(appsCmd)
}
