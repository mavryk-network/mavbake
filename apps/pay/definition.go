package pay

import (
	"github.com/mavryk-network/mavbake/apps/base"
)

func (app *Mavpay) LoadAppDefinition() (map[string]any, string, error) {
	return base.LoadAppDefinition(app)
}

func (app *Mavpay) LoadAppConfiguration() (map[string]any, error) {
	return base.LoadAppConfiguration(app)
}

func (app *Mavpay) GetActiveModel() (map[string]any, error) {
	return base.GetActiveModel(app)
}
