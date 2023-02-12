package upd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-ini/ini"
)

type Release struct {
	TagName string `json:"tag_name"`
}

func GetVersion() string {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
		return ""
	}
	return cfg.Section("data").Key("version").String()
}

func SaveVersion(v string) {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
		return
	}

	cfg.Section("data").Key("version").SetValue(v)

	err = cfg.SaveTo("config.ini")
	if err != nil {
		fmt.Printf("Error saving config file: %s\n", err)
		return
	}

}

func ChekVersion() string {
	url := "https://api.github.com/repos/PulsarG/PassManager/releases/latest"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error checking for updates:", err)
		return ""
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	}

	var release Release
	err = json.Unmarshal(body, &release)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return ""
	}

	return release.TagName
}
