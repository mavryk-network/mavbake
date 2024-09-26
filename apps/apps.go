package apps

import (
	"github.com/mavryk-network/mavbake/apps/base"
	"github.com/mavryk-network/mavbake/apps/node"
	"github.com/mavryk-network/mavbake/apps/pay"
	"github.com/mavryk-network/mavbake/apps/peak"
	"github.com/mavryk-network/mavbake/apps/signer"
)

var (
	Node   = node.FromPath("")
	Signer = signer.FromPath("")
	Peak   = peak.FromPath("")
	Pay    = pay.FromPath("")
	All    = []base.MavPayApp{
		Node, Signer, Peak, Pay,
	}
	Implicit = []base.MavPayApp{
		Node, Signer,
	}
)

type SetupContext = base.SetupContext
type UpgradeContext = base.UpgradeContext

type NodeInfoCollectionOptions = node.InfoCollectionOptions
type SignerInfoCollectionOptions = signer.InfoCollectionOptions

func GetInstalledApps() []base.MavPayApp {
	result := make([]base.MavPayApp, 0)
	for _, v := range All {
		if v.IsInstalled() {
			result = append(result, v)
		}
	}
	return result
}

func NodeFromPath(path string) *node.Node {
	return node.FromPath(path)
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
