package errloger

import (
	"PassManager/cons"

	"github.com/PulsarG/err-handler"
	"github.com/PulsarG/mailsend"

	"github.com/PulsarG/ConfigManager"
)

func ErrorLog(err error) {
	errorhandler.LoggError(err)
	if inihandler.GetFromIni("data", "sendlog") == "true" {
		mailsender.SendMail(cons.MAIL_FOR_ERROR, cons.MAIL_FOR_ERROR, cons.KEY_FOR_ERROR, err)
	}
}
