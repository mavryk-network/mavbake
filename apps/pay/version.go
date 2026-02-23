package pay

import (
	"github.com/mavryk-network/mavbake/ami"
)

func (app *Mavpay) GetVersions(options ami.CollectVersionsOptions) (*ami.InstanceVersions, error) {
	return ami.GetVersions(app.GetPath(), options)
}

func (app *Mavpay) GetVersion() (string, error) {
	return ami.GetAppVersion(app.GetPath())
}
