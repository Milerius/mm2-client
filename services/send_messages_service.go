package services

import (
	"context"
	"github.com/kpango/glg"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
	"mm2_client/config"
	"mm2_client/constants"
	"time"
)

var gNotifier *notify.Notify = nil

func init() {
	gNotifier = notify.New()
}

func loadNotifierCFG(kind string, target string) {
	switch kind {
	case "file":
		if config.ParseNotifyCfgFromFile(target) {
			_ = glg.Infof("Notify cfg: %s successfully loaded", target)
		}
	default:
		_ = glg.Errorf("%s - not supported", kind)
	}
}

func SendMessage(subject string, message string) {
	if constants.GNotifyCfgLoaded {
		_ = gNotifier.Send(
			context.Background(),
			subject,
			message,
		)
	} else {
		_ = glg.Warn("You try to send a message to the notify service, but the configuration is not loaded - skipping")
	}
}

func StartNotifierMessagesService(kind string, target string) {
	_ = glg.Infof("CreateNotifierMessagesService")
	loadNotifierCFG(kind, target)
	if config.GNotifyCFG != nil {
		if config.GNotifyCFG.Telegram != nil {
			telegramService, telegramErr := telegram.New(config.GNotifyCFG.Telegram.TelegramApiToken)
			if telegramErr == nil {
				telegramService.AddReceivers(config.GNotifyCFG.Telegram.TelegramReceiver)
				gNotifier.UseServices(telegramService)
				_ = glg.Info("Successfully added telegram service")
				SendMessage(time.Now().String()+" - Milerius LP Bot", "Starting Notify Service for LP Provider")
			} else {
				_ = glg.Errorf("telegram error: %v", telegramErr)
			}
		}
	}
}
