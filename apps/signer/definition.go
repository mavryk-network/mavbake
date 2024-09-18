package signer

import (
	"github.com/mavryk-network/mavbake/apps/base"
)

func (app *Signer) LoadAppDefinition() (map[string]interface{}, string, error) {
	return base.LoadAppDefinition(app)
}
func (app *Signer) LoadAppConfiguration() (map[string]interface{}, error) {
	return base.LoadAppConfiguration(app)
}
