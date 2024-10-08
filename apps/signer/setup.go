package signer

import (
	"fmt"

	"github.com/mavryk-network/mavbake/ami"
	"github.com/mavryk-network/mavbake/apps/base"
	"github.com/mavryk-network/mavbake/constants"
	"github.com/mavryk-network/mavbake/util"

	log "github.com/sirupsen/logrus"
)

func (app *Signer) GetSetupKind() string {
	return base.MergingSetupKind
}

func (app *Signer) Setup(ctx *base.SetupContext, args ...string) (int, error) {
	appDef, err := base.GenerateConfiguration(app.GetAmiTemplate(ctx), ctx)
	if err != nil {
		log.Warn(err)
	}

	oldAppDef, err := ami.ReadAppDefinition(app.GetPath(), constants.DefaultAppJsonName)
	if oldAppDef != nil && err == nil {
		if oldConfiguration, ok := (*oldAppDef)["configuration"].(map[string]interface{}); ok {
			log.Info("Found old configuration. Merging...")
			appDef["configuration"] = util.MergeMaps(oldConfiguration, appDef["configuration"].(map[string]interface{}), true)
		}
	}

	err = ami.WriteAppDefinition(app.GetPath(), appDef, constants.DefaultAppJsonName)
	if err != nil {
		return -1, fmt.Errorf("failed to write app definition - %s", err.Error())
	}
	return ami.SetupApp(app.GetPath(), args...)
}
