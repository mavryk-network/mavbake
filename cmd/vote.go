package cmd

import (
	"fmt"

	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/constants"
	"github.com/mavryk-network/mavbake/util"

	"github.com/spf13/cobra"
)

var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "(DEPRECATED) Cast vote for proposal/s.",
	Long: `(DEPRECATED) Use 'https://gov.tez.capital/' instead.
	
	Submits vote for proposal/s based on provided voting period.
	
	Exploration:
	` + "`" + `mavbake vote --period exploration <proposal> yay|nay|pass` + "`" + `
	Proposal:
	` + "`" + `mavbake vote --period proposal <proposal1> <proposal2>` + "`" + `
	Promotion:
	` + "`" + `mavbake vote --period promotion <proposal> yay|nay|pass` + "`" + `
	`,
	Run: func(cmd *cobra.Command, args []string) {
		period, _ := cmd.Flags().GetString("period")
		voteArgs := make([]string, 0)
		voteArgs = append(voteArgs, "client")
		// TODO:
		// if period == "auto" {
		// 	bb.Node.GetInfo()
		// }
		switch period {
		case "proposal":
			voteArgs = append(voteArgs, "submit", "proposals", "for", "baker")
			voteArgs = append(voteArgs, args...)

			_, err := apps.Signer.Execute(voteArgs...)
			util.AssertE(err, "Failed to vote in '"+period+"' for "+fmt.Sprintf("%v", args)+"!")
		case "exploration":
			voteArgs = append(voteArgs, "submit", "ballot", "for", "baker")
			voteArgs = append(voteArgs, args...)

			_, err := apps.Signer.Execute(voteArgs...)
			util.AssertE(err, "Failed to vote in '"+period+"' for "+fmt.Sprintf("%v", args)+"!")
		case "promotion":
			//mavryk-client submit ballot for YOUR_ADDRESS Psithaca2MLRFYargivpo7YvUr7wUDqyxrdhC5CQq78mRvimz6A yay
			voteArgs = append(voteArgs, "submit", "ballot", "for", "baker")
			voteArgs = append(voteArgs, args...)

			_, err := apps.Signer.Execute(voteArgs...)
			util.AssertE(err, "Failed to vote in '"+period+"' for "+fmt.Sprintf("%v", args)+"!")
		default:
			util.AssertBE(false, "Invalid period - '"+period+"'!", constants.ExitInvalidArgs)
		}
	},
}

func init() {
	voteCmd.Flags().String("period", "unknown", "Sets period to vote on.")
	voteCmd.Flags().SetInterspersed(false)

	RootCmd.AddCommand(voteCmd)
}
