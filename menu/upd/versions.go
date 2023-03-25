// CHECKING AND OBTAINING A NEW VERSION NUMBER

package upd

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("Error checking for updates:", err)
		return ""
	} // end if
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	} // end if

	var release Release
	err = json.Unmarshal(body, &release)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return ""
	} // end if

	return release.TagName
}
