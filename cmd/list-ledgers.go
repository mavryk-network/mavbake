package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"

	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/cli"
	"github.com/mavryk-network/mavbake/constants"
	"github.com/mavryk-network/mavbake/util"
	"go.alis.is/common/log"

	"github.com/spf13/cobra"
)

var listLedgersCmd = &cobra.Command{
	Use:   "list-ledgers",
	Short: "Prints list of available ledgers.",
	Long:  "Collects and prits list of available ledger ids.",
	Run: func(cmd *cobra.Command, args []string) {
		mavClientPath := path.Join(apps.Signer.GetPath(), "bin", "client")
		log.Trace("Listing connected ledgers:", "mav_client_path", mavClientPath)
		output, err := exec.Command(mavClientPath, "list", "connected", "ledgers").CombinedOutput()
		if matched, _ := regexp.Match("Error:", output); err != nil || matched {
			fmt.Println(string(output))
			log.Error("Failed to list ledgers!", "error", err)
			os.Exit(constants.ExitExternalError)
		}
		matchLedgers := regexp.MustCompile("## Ledger `(.*?)`")
		matches := matchLedgers.FindAllStringSubmatch(string(output), -1)
		if cli.JsonLogFormat {
			res := make([]string, 0)
			for _, v := range matches {
				if len(v) > 1 {
					res = append(res, v[1])
				}
			}
			output, err := json.Marshal(res)
			util.AssertEE(err, "Failed to serialize list of ledgers!", constants.ExitSerializationFailed)
			fmt.Println(string(output))
		} else {
			for _, v := range matches {
				if len(v) > 1 {
					fmt.Println(v[1])
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listLedgersCmd)
}
