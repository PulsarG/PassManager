package upd

import (
	"PassManager/confile"
	"PassManager/cons"
	"net/http"
	"os"
	"os/exec"

	"github.com/inconshreveable/go-update"
)

func Update() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	newVersion := ChekVersion()
	url := cons.URL_FOR_DOWNLOAD + newVersion + "/auto-update-file.exe "

	resp, err := http.Get(url)
	if err != nil {
		return "Fail download file"
	}
	defer resp.Body.Close()

	oldPath := update.Apply(resp.Body, update.Options{
		TargetPath: exePath,
	})
	if err != nil {
		if rerr := update.RollbackError(err); rerr != nil {
			return "Fail update file"
		}
	}

	confile.SaveToIni("data", "version", newVersion)
	confile.SaveToIni("data", "old", oldPath)

	executablePath, err := os.Executable()
	if err != nil {
		return "Other Fail"
	}
	err = exec.Command(executablePath).Start()
	if err != nil {
		return "Fail start new version"
	}
	
	os.Exit(0)

	return "Oops"
}
