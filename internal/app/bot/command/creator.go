package command

import (
	"PriceWatcher/internal/entities/bot"
	"PriceWatcher/internal/entities/subscribing"
	"sync"
)

func CreateSubCommand(mu *sync.Mutex, subs *subscribing.Subscribers) bot.Command {
	subCom := newSubCommand(mu, subs)
	return bot.Command{
		Name:        "start",
		Description: "Start getting messages of the current gold price ",
		Action:      subCom.subscribeUser,
	}
}

func CreateUnsubCommand(mu *sync.Mutex, subs *subscribing.Subscribers) bot.Command {
	unsubCom := newUnsubCommand(mu, subs)
	return bot.Command{
		Name:        "stop",
		Description: "Stop getting notifications about the current gold price ",
		Action:      unsubCom.unsubscribeUser,
	}
}
