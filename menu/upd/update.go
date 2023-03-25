// UPDATE APP. DOWNLOAD NEW VERSION EXE AND START RUN THIS

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
	} // end if

	newVersion := ChekVersion()
	url := cons.URL_FOR_DOWNLOAD + newVersion + "/auto-update-file.exe "

	resp, err := http.Get(url)
	if err != nil {
		return "Fail download file"
	} // end if
	defer resp.Body.Close()

	oldPath := update.Apply(resp.Body, update.Options{
		TargetPath: exePath,
	})
	if err != nil { // if outer
		if rerr := update.RollbackError(err); rerr != nil { // if inner
			return "Fail update file"
		} // end if inner
	} // end if outer

	confile.SaveToIni("data", "version", newVersion)
	confile.SaveToIni("data", "old", oldPath)

	executablePath, err := os.Executable()
	if err != nil {
		return "Other Fail"
	} // end if
	err = exec.Command(executablePath).Start()
	if err != nil {
		return "Fail start new version"
	} // end if
	
	os.Exit(0)

	return "Oops"
}
