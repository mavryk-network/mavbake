package constants

import (
	"runtime"
)

const (
	MavbakeRepository string = "mavryk-network/mavbake"

	defaultBBDirectory      string = "/mavbake"
	defaultBBDirectoryMacOS string = "/usr/local/mavbake"
	DefaultRemoteUser       string = "mavbake"
	DefaultSshUser          string = "root"

	DefaultAppJsonName string = "app.json"

	MvktConsensusKeyCheckingEndpoint = "https://api.mavryk.network/"
)

var (
	DefaultBBDirectory string = defaultBBDirectory
)

func init() {
	if runtime.GOOS == "darwin" {
		DefaultBBDirectory = defaultBBDirectoryMacOS
	}
}
