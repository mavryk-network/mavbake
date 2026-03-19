package cmd

import (
	"fmt"
	"os"

	"github.com/mavryk-network/mavbake/ami"
	"github.com/mavryk-network/mavbake/apps"
	"github.com/mavryk-network/mavbake/constants"
	"github.com/mavryk-network/mavbake/system"
	"github.com/mavryk-network/mavbake/util"
	"go.alis.is/common/log"

	"github.com/spf13/cobra"
)

var (
	mavsignPlatformFlag  *BoolStringCombinedFlag
	mavsignImportKeyFlag *BoolStringCombinedFlag
)

var setupMavsignCmd = &cobra.Command{
	Use:    "setup-mavsign",
	Short:  "Setup mavsign for baking.",
	Long:   "Setups mavsign for baking.",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		shouldOperateOnSigner, _ := cmd.Flags().GetBool("signer")
		shouldOperateOnNode, _ := cmd.Flags().GetBool("node")
		init, _ := cmd.Flags().GetBool("init")
		force, _ := cmd.Flags().GetBool("force")
		keyAlias, _ := cmd.Flags().GetString("key-alias")
		password, _ := cmd.Flags().GetBool("password")

		isAnySelected := shouldOperateOnSigner || shouldOperateOnNode

		if mavsignImportKeyFlag.IsTrue() && (mavsignPlatformFlag.IsTrue() || init) {
			log.Error("Cannot use --import-key together with --platform or --init. Please run setup-mavsign in two steps.")
			os.Exit(constants.ExitInvalidArgs)
			return
		}

		if mavsignPlatformFlag.IsTrue() || init || password {
			system.RequireElevatedUser() // platform setup, init and password require elevated permissions
		}

		if mavsignImportKeyFlag.IsTrue() { // mavsign import requires signer to be running
			log.Info("ensuring signer is running for mavsign key import...")
			wasRunning, _ := apps.Signer.IsServiceStatus(constants.MavpayAppServiceId, "running")
			if !wasRunning {
				system.RequireElevatedUser() // starting signer service requires elevated permissions
				exitCode, err := apps.Signer.Start()
				util.AssertEE(err, "Failed to start signer!", exitCode)
				defer apps.Signer.Stop()
			}
		}

		if (shouldOperateOnSigner || !isAnySelected) && apps.Signer.IsInstalled() {
			log.Info("setting up mavsign for signer...")

			amiArgs := []string{"setup-mavsign"}

			if init {
				amiArgs = append(amiArgs, "--init")
			}
			if mavsignPlatformFlag.HasValue() {
				amiArgs = append(amiArgs, "--platform="+mavsignPlatformFlag.String())
			} else if mavsignPlatformFlag.IsTrue() {
				amiArgs = append(amiArgs, "--platform")
			}

			noUdev, _ := cmd.Flags().GetString("no-udev")
			if noUdev != "" {
				amiArgs = append(amiArgs, "--no-udev")
			}

			if password {
				amiArgs = append(amiArgs, "--password")
			}

			if mavsignImportKeyFlag.HasValue() {
				amiArgs = append(amiArgs, "--import-key="+mavsignImportKeyFlag.String())
			} else if mavsignImportKeyFlag.IsTrue() {
				amiArgs = append(amiArgs, "--import-key")
			}

			amiArgs = append(amiArgs, fmt.Sprintf("--key-alias=%s", keyAlias))

			if force {
				amiArgs = append(amiArgs, "--force")
			}

			exitCode, err := apps.Signer.Execute(amiArgs...)
			util.AssertEE(err, "Failed to import key to signer!", exitCode)
		}

		if (shouldOperateOnNode || !isAnySelected) && apps.Node.IsInstalled() && mavsignImportKeyFlag.IsTrue() { // node only imports key
			log.Info("Importing key to the node...")

			isSignerRunning, _ := apps.Signer.IsServiceStatus(constants.SignerAppServiceId, "running")
			util.AssertBE(isSignerRunning, "Signer is not running. Please start signer services.", constants.ExitSignerNotOperational)

			bakerAddr, exitCode, err := apps.Signer.GetKeyHash(keyAlias)
			util.AssertEE(err, "Failed to get baker key hash!", exitCode)
			ami.REMOTE_VARS[ami.BAKER_KEY_HASH_REMOTE_VAR] = bakerAddr
			amiArgs := []string{"import-key", bakerAddr}
			if force {
				amiArgs = append(amiArgs, "--force")
			}
			amiArgs = append(amiArgs, fmt.Sprintf("--alias=%s", keyAlias))
			exitCode, err = apps.Node.Execute(amiArgs...)
			util.AssertEE(err, "Failed to import key to node!", exitCode)

		}
	},
}

func init() {
	setupMavsignCmd.Flags().Bool("node", false, "Import key to node (affects import-key only)")
	setupMavsignCmd.Flags().Bool("signer", false, "Import key to signer (affects import-key only)")
	setupMavsignCmd.Flags().Bool("init", false, "Initialize mavsign configuration.")
	setupMavsignCmd.Flags().Bool("password", false, "Setup mavsign unlock password.")

	mavsignImportKeyFlag = addCombinedFlag(setupMavsignCmd, "import-key", "", "Import key from mavsign (optionally specify derivation path)")
	setupMavsignCmd.Flags().String("key-alias", "baker", "Alias ofkey to be imported")

	mavsignPlatformFlag = addCombinedFlag(setupMavsignCmd, "platform", "", "Prepare platform for mavsign (optionally specify platform to override)")
	setupMavsignCmd.Flags().String("no-udev", "", "Skip udev rules installation. (linux only)")

	setupMavsignCmd.Flags().BoolP("force", "f", false, "Force key import. (overwrites existing)")

	RootCmd.AddCommand(setupMavsignCmd)
}
