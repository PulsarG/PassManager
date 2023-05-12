// CHECKING AND OBTAINING A NEW VERSION NUMBER

package upd

import (
	// "PassManager/confile"
	"PassManager/errloger"

	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"

	"PassManager/cons"
)

type Release struct {
	TagName string `json:"tag_name"`
}

func ChekVersion() string {
	response, err := http.Get(cons.URL_LATEST_VERSION)
	if err != nil {
		errloger.ErrorLog(err)
		return ""
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errloger.ErrorLog(err)
		return ""
	}

	var release Release
	err = json.Unmarshal(body, &release)
	if err != nil {
		errloger.ErrorLog(err)
		return ""
	}

	return release.TagName
}
