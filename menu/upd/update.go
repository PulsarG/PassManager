// UPDATE APP. DOWNLOAD NEW VERSION EXE AND START RUN THIS

package upd

import (
	// "PassManager/errs"
	"PassManager/errloger"
	// "PassManager/confile"
	"PassManager/cons"
	"net/http"
	"os"
	"os/exec"

	"github.com/inconshreveable/go-update"

	"github.com/PulsarG/ConfigManager"
)

func Update() string {
	exePath, err := os.Executable()
	if err != nil {
		errloger.ErrorLog(err)
		panic(err)
	}

	newVersion := ChekVersion()
	url := cons.URL_FOR_DOWNLOAD + newVersion + "/auto-update-file.exe "

	resp, err := http.Get(url)
	if err != nil {
		errloger.ErrorLog(err)
		return "Fail download file"
	}
	defer resp.Body.Close()

	oldPath := update.Apply(resp.Body, update.Options{
		TargetPath: exePath,
	})
	if err != nil {
		if rerr := update.RollbackError(err); rerr != nil {
			errloger.ErrorLog(err)
			return "Fail update file"
		}
	}

	inihandler.SaveToIni("data", "version", newVersion)
	inihandler.SaveToIni("data", "old", oldPath)

	executablePath, err := os.Executable()
	if err != nil {
		errloger.ErrorLog(err)
		return "Other Fail"
	}
	err = exec.Command(executablePath).Start()
	if err != nil {
		errloger.ErrorLog(err)
		return "Fail start new version"
	}

	os.Exit(0)

	return "Oops"
}
