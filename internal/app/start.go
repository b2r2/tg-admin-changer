package app

import (
	"github.com/b2r2/tg-admin-changer/internal/config"
	"github.com/b2r2/tg-admin-changer/internal/models"

	tele "gopkg.in/tucnak/telebot.v3"
)

func (b *bot) OnStart() tele.HandlerFunc {
	return func(c tele.Context) error {
		var msg = config.Get().GetMapping()[models.Greeting]
		opts := &tele.SendOptions{ReplyMarkup: b.inlineMainMenu}
		m, err := b.bot.Send(&tele.Chat{ID: c.Chat().ID}, msg.String(), opts)
		if err != nil {
			b.log.Println("OnStart(send message)", err)
		}
		if err = b.bot.Pin(m); err != nil {
			b.log.Println("OnStart(pin message)", err)
		}

		return nil
	}
}
