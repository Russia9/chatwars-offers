package app

import (
	"gitea.russia9.dev/Russia9/chatwars-offers/messages"
	"github.com/rs/zerolog/log"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
)

func (a *App) Sender(channel chan messages.OfferMessage) {
	for {
		var message messages.OfferMessage
		message = <-channel

		msgString :=
			" " + message.SellerCastle + message.SellerName + ": \n" +
				" " + strconv.Itoa(message.Quantity) + " " + message.Item + " *💰" + strconv.Itoa(message.Price)

		_, err := a.Bot.Send(a.Chat, msgString, telebot.ModeHTML)
		if err != nil {
			log.Error().Err(err).Str("module", "sender").Send()
		}
	}
}
