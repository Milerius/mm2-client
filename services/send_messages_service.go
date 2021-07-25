package services

import (
	"github.com/kpango/glg"
	"github.com/nikoksr/notify"
)

var gNotifier *notify.Notify = nil

func init() {
	gNotifier = notify.New()
}

func StartNotifierMessagesService() {
	glg.Infof("CreateNotifierMessagesService")
}
