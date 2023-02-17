package upd

import (
	"fmt"
	"github.com/inconshreveable/go-update"
	"net/http"
)

func Update(){
	// Получаем URL для загрузки новой версии
	newVersion := ChekVersion()
	url := "https://github.com/PulsarG/PassManager/releases/download/" + newVersion + "/EnigmaPass_" + newVersion + "_Win10.7z"

	// Отправляем запрос на URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading update:", err)
		return
	}
	defer resp.Body.Close()

	// Обновляем приложение
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		fmt.Println("Error updating application:", err)
		return
	}

	fmt.Println("Application updated successfully.")
}
