package peak

import (
	"github.com/mavryk-network/mavbake/ami"

	"github.com/mavryk-network/mavbake/apps/base"
)

func (app *Peak) Upgrade(ctx *base.UpgradeContext, args ...string) (int, error) {
	wasRunning, _ := app.IsServiceStatus("mavpeak", "running")
	if wasRunning {
		exitcode, err := app.Stop()
		if err != nil {
			return exitcode, err
		}
	}
	exitCode, err := ami.SetupApp(app.GetPath(), args...)
	if wasRunning {
		exitcode, err := app.Start()
		if err != nil {
			return exitcode, err
		}
	}
	return exitCode, err
}
