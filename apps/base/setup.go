package base

import (
	"path"

	"github.com/mavryk-network/mavbake/ami"
	"github.com/mavryk-network/mavbake/constants"
	"github.com/mavryk-network/mavbake/system"
)

type SetupContext struct {
	Force bool
	// general info
	Configuration string
	Version       string
	User          string
	Branch        string

	// remote info
	Remote                string
	RemoteAuth            string
	RemoteElevate         ami.RemoteElevationKind
	RemoteElevatePassword string
	RemoteReset           bool

	Dal bool
}

func (ctx *SetupContext) ToRemoteConfiguration(app MavBakeApp) *ami.RemoteConfiguration {
	connectionDetails := system.GetRemoteConnectionDetails(ctx.Remote)

	return &ami.RemoteConfiguration{
		ElevationCredentialsDirectory: app.GetPath(),
		App:                           app.GetId(),
		Username:                      connectionDetails.Username,
		LocalUsername:                 ctx.User,
		Host:                          connectionDetails.Host,
		Port:                          connectionDetails.Port,
		InstancePath:                  constants.DefaultBBDirectory,
		Elevate:                       ctx.RemoteElevate,
		PrivateKey:                    path.Join(app.GetPath(), constants.PrivateKeyFile),
		PublicKey:                     path.Join(app.GetPath(), constants.PublicKeyFile),
	}
}

func (ctx *SetupContext) ToRemoteElevateCredentials() *ami.RemoteElevateCredentials {
	if ctx.RemoteElevate == ami.REMOTE_ELEVATION_NONE {
		return nil
	}
	return &ami.RemoteElevateCredentials{
		Kind: ami.RemoteElevationKind(ctx.RemoteElevate),
		// User:     ctx.RemoteElevateUser,
		Password: ctx.RemoteElevatePassword,
	}
}
