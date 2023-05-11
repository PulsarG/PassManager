package confile

import (
	// "PassManager/confile"
	"PassManager/cons"
	// "fmt"
	"github.com/PulsarG/err-handler"
	"github.com/PulsarG/mailsend"
	// "net/smtp"
	// "strings"

	"github.com/PulsarG/ConfigManager"
)

func ErrorLog(err error) {
	isLocal := GetFromIni("data", "loglocal")
	if isLocal == "true" {
		errorhandler.LoggError(err)
	} else {
		// fmt.Println(err.Error())
		mailsender.SendMail(cons.MAIL_FOR_ERROR, cons.MAIL_FOR_ERROR, cons.KEY_FOR_ERROR, err)
	}

	inihandler.GetFromIni()
}
