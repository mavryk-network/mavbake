package pay

import (
	"github.com/mavryk-network/mavbake/apps/base"
)

func (app *Mavpay) LoadAppDefinition() (map[string]interface{}, string, error) {
	return base.LoadAppDefinition(app)
}
func (app *Mavpay) LoadAppConfiguration() (map[string]interface{}, error) {
	return base.LoadAppConfiguration(app)
}
