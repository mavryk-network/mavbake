package peak

import (
	"github.com/mavryk-network/mavbake/ami"
)

func (app *Peak) GetVersions(options ami.CollectVersionsOptions) (*ami.InstanceVersions, error) {
	return ami.GetVersions(app.GetPath(), options)
}

func (app *Peak) GetVersion() (string, error) {
	return ami.GetAppVersion(app.GetPath())
}
