package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mavryk-network/mavbake/ami"
	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/apps/base"
	"github.com/mavryk-network/mavbake/cli"
	"github.com/mavryk-network/mavbake/constants"
	"github.com/mavryk-network/mavbake/util"

	"github.com/jedib0t/go-pretty/v6/table"
	lop "github.com/samber/lo/parallel"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints MavBake CLI version.",
	Long:  "Prints MavBake CLI version.",
	Run: func(cmd *cobra.Command, args []string) {
		shouldPrintAll, _ := cmd.Flags().GetBool("all")
		// shouldPrintPackages, _ := cmd.Flags().GetBool("packages")
		// shouldPrintBinaries, _ := cmd.Flags().GetBool("binaries")

		// shouldPrintPackages = shouldPrintPackages || !shouldPrintBinaries || shouldPrintAll
		// shouldPrintBinaries = shouldPrintBinaries || !shouldPrintPackages || shouldPrintAll

		appsToCollectFrom := GetAppsBySelectionCriteria(cmd, AppSelectionCriteria{
			InitialSelection:  InstalledApps,
			FallbackSelection: NoFallback,
		})

		if len(appsToCollectFrom) == 0 {
			if cli.JsonLogFormat {
				ver, _ := json.Marshal(constants.VERSION)
				fmt.Print(string(ver))
				return
			}
			fmt.Println(constants.VERSION)
			return
		}

		if shouldPrintAll && len(appsToCollectFrom) == 0 {
			appsToCollectFrom = apps.GetInstalledApps(cmd)
		}

		versionTable := table.NewWriter()
		versionTable.SetOutputMirror(os.Stdout)
		versionTable.SetStyle(table.StyleLight)
		versionTable.AppendHeader(table.Row{"Versions", "Versions"}, table.RowConfig{AutoMerge: true})
		versionTable.AppendRow(table.Row{"mavbake", constants.VERSION})

		switch {
		case shouldPrintAll:

			appVersions := lop.Map(appsToCollectFrom, func(v base.MavBakeApp, _ int) *ami.InstanceVersions {
				versions, err := v.GetVersions(ami.CollectVersionsOptions{})
				util.AssertE(err, fmt.Sprintf("Failed to collect %s's versions!", v.GetId()))
				return versions
			})
			switch {
			case cli.JsonLogFormat:
				result := make(map[string]any)
				result["mavbake"] = constants.VERSION
				for i, v := range appsToCollectFrom {
					versions := appVersions[i]
					result[v.GetId()] = versions
				}
				verInfo, err := json.Marshal(result)
				util.AssertEE(err, "Failed to serialize version info!", constants.ExitSerializationFailed)
				fmt.Print(string(verInfo))
				return
			default:
				for i, v := range appsToCollectFrom {
					versions := appVersions[i]
					versionTable.AppendSeparator()
					versionTable.AppendRow(table.Row{v.GetLabel(), v.GetLabel()}, table.RowConfig{AutoMerge: true})
					versionTable.AppendSeparator()
					if versions.RemoteMavbake != "" {
						versionTable.AppendRow(table.Row{"mavbake (remote)", versions.RemoteMavbake}, table.RowConfig{AutoMerge: true})
						versionTable.AppendSeparator()
					}
					if len(versions.Packages) > 0 {
						versionTable.AppendRow(table.Row{"Packages", "Packages"}, table.RowConfig{AutoMerge: true})
						versionTable.AppendSeparator()
						for k, v := range versions.Packages {
							versionTable.AppendRow(table.Row{k, v})
						}
						versionTable.AppendSeparator()
					}
					if len(versions.Binaries) > 0 {
						versionTable.AppendRow(table.Row{"Binaries", "Binaries"}, table.RowConfig{AutoMerge: true})
						versionTable.AppendSeparator()
						for k, v := range versions.Binaries {
							versionTable.AppendRow(table.Row{k, v})
						}
					}
				}
			}

		default:
			appVersions := lop.Map(appsToCollectFrom, func(v base.MavBakeApp, _ int) string {
				version, err := v.GetVersion()
				util.AssertE(err, fmt.Sprintf("Failed to collect %s's versions!", v.GetId()))
				return version
			})

			switch {
			case cli.JsonLogFormat:
				result := make(map[string]any)
				result["mavbake"] = constants.VERSION
				for i, v := range appsToCollectFrom {
					version := appVersions[i]
					result[v.GetId()] = version
				}
				verInfo, err := json.Marshal(result)
				util.AssertEE(err, "Failed to serialize version info!", constants.ExitSerializationFailed)
				fmt.Print(string(verInfo))
				return
			default:
				for i, v := range appsToCollectFrom {
					version := appVersions[i]
					versionTable.AppendRow(table.Row{v.GetLabel(), version}, table.RowConfig{AutoMerge: false})
				}
			}
		}

		versionTable.Render()
	},
}

func init() {
	versionCmd.Flags().BoolP("all", "a", false, "Prints version of all MavBake instance packages/binaries.")
	for _, v := range apps.All {
		versionCmd.Flags().Bool(v.GetId(), false, fmt.Sprintf("Prints versions of %s.", v.GetId()))
	}
	// versionCmd.Flags().Bool("packages", false, "Prints versions packages.")
	// versionCmd.Flags().BoolP("binaries", "b", false, "Prints versions of binaries.")
	RootCmd.AddCommand(versionCmd)
}
