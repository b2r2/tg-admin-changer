package app

import (
	"github.com/b2r2/tg-admin-changer/internal/config"
	"github.com/b2r2/tg-admin-changer/internal/models"

	tele "gopkg.in/tucnak/telebot.v3"
)

func (b *bot) OnStart() tele.HandlerFunc {
	return func(c tele.Context) error {
		var msg = config.Get().GetMapping()[models.Greeting]
		_, err := b.bot.Send(&tele.Chat{ID: c.Chat().ID}, msg.String(), &tele.SendOptions{
			ReplyMarkup: menu,
		})
		if err != nil {
			b.log.Println("OnStart(send message)", err)
		}
		return nil
	}
}
