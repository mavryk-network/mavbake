package peak

import (
	"path"
	"strings"

	"github.com/mavryk-network/mavbake/ami"
	"github.com/mavryk-network/mavbake/apps/base"
	"github.com/mavryk-network/mavbake/cli"
	"github.com/mavryk-network/mavbake/constants"
)

var (
	Id           string                 = constants.PeakAppId
	AMI_TEMPLATE map[string]interface{} = map[string]interface{}{
		"id":            constants.PeakAppId,
		"type":          map[string]interface{}{"id": "mvd.mavpeak", "version": "latest"},
		"configuration": map[string]interface{}{},
		"user":          "",
	}
)

type Peak struct {
	Path string
}

// FromPath creates a new peak instance with the specified path.
// The path parameter is the directory path to be associated with the peak.
// If the path is empty, the default path will be used.
// It returns a pointer to the newly created peak instance.
func FromPath(path string) *Peak {
	return &Peak{
		Path: path,
	}
}

func (app *Peak) GetAmiTemplate(ctx *base.SetupContext) map[string]interface{} {
	return AMI_TEMPLATE
}

func (app *Peak) GetPath() string {
	if app.Path != "" {
		return app.Path
	}
	return path.Join(cli.BBdir, Id)
}

func (app *Peak) GetId() string {
	return strings.ToLower(constants.PeakAppId)
}

func (app *Peak) GetLabel() string {
	return strings.ToUpper(app.GetId())
}

func (app *Peak) IsInstalled() bool {
	return ami.IsAppInstalled(app.GetPath())
}

func (app *Peak) SupportsRemote() bool {
	return false
}

func (app *Peak) IsRemoteApp() bool {
	return false
}
