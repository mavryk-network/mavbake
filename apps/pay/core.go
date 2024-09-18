package pay

import (
	"path"
	"strings"

	"github.com/mavryk-network/mavbake/ami"
	"github.com/mavryk-network/mavbake/apps/base"
	"github.com/mavryk-network/mavbake/cli"
	"github.com/mavryk-network/mavbake/constants"
)

var (
	Id           string                 = constants.MavpayAppId
	AMI_TEMPLATE map[string]interface{} = map[string]interface{}{
		"id":            constants.MavpayAppId,
		"type":          map[string]interface{}{"id": "tzc.mavpay", "version": "latest"},
		"configuration": map[string]interface{}{},
		"user":          "",
	}
)

type Mavpay struct {
	Path string
}

// FromPath creates a new Node instance with the specified path.
// The path parameter is the directory path to be associated with the Node.
// If the path is empty, the default path will be used.
// It returns a pointer to the newly created Node instance.
func FromPath(path string) *Mavpay {
	return &Mavpay{
		Path: path,
	}
}

func (app *Mavpay) GetPath() string {
	if app.Path != "" {
		return app.Path
	}
	return path.Join(cli.BBdir, Id)
}

func (app *Mavpay) GetId() string {
	return strings.ToLower(constants.MavpayAppId)
}

func (app *Mavpay) GetLabel() string {
	return strings.ToUpper(app.GetId())
}

func (app *Mavpay) GetAmiTemplate(ctx *base.SetupContext) map[string]interface{} {
	return AMI_TEMPLATE
}
func (app *Mavpay) IsInstalled() bool {
	return ami.IsAppInstalled(app.GetPath())
}

func (app *Mavpay) SupportsRemote() bool {
	return false
}

func (app *Mavpay) IsRemoteApp() bool {
	return false
}
