package ami

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mavryk-network/mavbake/util"

	"path"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var (
	amiInstallScriptSource = "https://raw.githubusercontent.com/alis-is/ami/master/install.sh"
)

func Install() (int, error) {
	log.Trace("Downloading eli&ami install script...")

	tmpInstallScript := path.Join(os.TempDir(), fmt.Sprintf("%s-%s", uuid.NewString(), "install.sh"))
	err := util.DownloadFile(amiInstallScriptSource, tmpInstallScript, false)
	if err != nil {
		return -1, err
	}
	defer os.Remove(tmpInstallScript)

	shPath, err := exec.LookPath("sh")
	if err != nil {
		return -1, err
	}
	log.Trace("Executing eli&ami install script...")
	installProc := exec.Command(shPath, tmpInstallScript)
	installProc.Stdout = os.Stdout
	installProc.Stderr = os.Stderr
	err = installProc.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode(), err
		}
		return -1, err
	}
	return 0, nil
}
