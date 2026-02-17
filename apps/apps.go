package apps

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/mavryk-network/mavbake/apps/base"
	"github.com/mavryk-network/mavbake/apps/dal"
	"github.com/mavryk-network/mavbake/apps/node"
	"github.com/mavryk-network/mavbake/apps/pay"
	"github.com/mavryk-network/mavbake/apps/peak"
	"github.com/mavryk-network/mavbake/apps/signer"
)

var (
	Node    = node.FromPath("")
	DalNode = dal.FromPath("")
	Signer  = signer.FromPath("")
	Peak    = peak.FromPath("")
	Pay     = pay.FromPath("")
	All     = []base.MavBakeApp{
		Node, Signer, DalNode, Peak, Pay,
	}
	Implicit = []base.MavBakeApp{
		Node, Signer,
	}
)

type SetupContext = base.SetupContext
type UpgradeContext = base.UpgradeContext

type NodeInfoCollectionOptions = node.InfoCollectionOptions
type DalNodeInfoCollectionOptions = dal.InfoCollectionOptions
type SignerInfoCollectionOptions = signer.InfoCollectionOptions

func GetInstalledApps(cmd *cobra.Command) []base.MavBakeApp {
	result := make([]base.MavBakeApp, 0)
	initial := All
	filteredAll := lo.Filter(initial, func(app base.MavBakeApp, _ int) bool {
		found, _ := cmd.Flags().GetBool(app.GetId())
		return found
	})
	if len(filteredAll) > 0 {
		initial = filteredAll
	}
	for _, v := range initial {
		if v.IsInstalled() {
			result = append(result, v)
		}
	}
	return result
}

func NodeFromPath(path string) *node.Node {
	return node.FromPath(path)
}

func DalNodeFromPath(path string) *dal.DalNode {
	return dal.FromPath(path)
}

func SignerFromPath(path string) *signer.Signer {
	return signer.FromPath(path)
}

func PeakFromPath(path string) *peak.Peak {
	return peak.FromPath(path)
}

func MavpayFromPath(path string) *pay.Mavpay {
	return pay.FromPath(path)
}
