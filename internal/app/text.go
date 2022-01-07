package app

import (
	"github.com/b2r2/tg-admin-changer/internal/models"
	tele "gopkg.in/tucnak/telebot.v3"
)

func (b *bot) OnText() tele.HandlerFunc {
	return func(c tele.Context) error {
		cid := c.Sender().ID
		msg := c.Message()

		if _, ok := b.admins[cid]; !ok {
			if _, err := b.bot.Forward(&tele.Chat{ID: models.Channel}, msg); err != nil {
				b.log.Println("OnText(forward message)", err)
			}
			return nil
		}

		if msg.ReplyTo == nil {
			if _, err := b.bot.Send(&tele.Chat{ID: models.Channel}, models.NegativeMessage); err != nil {
				b.log.Println("OnText(forward message)", err)
			}
			return nil
		}

		if msg.ReplyTo.OriginalSender == nil {
			if _, err := b.bot.Send(&tele.Chat{ID: models.Channel}, models.NegativeMessageToBot); err != nil {
				b.log.Println("OnText(forward message)", err)
			}
			return nil
		}

		cid = msg.ReplyTo.OriginalSender.ID
		if _, err := b.bot.Send(&tele.Chat{ID: cid}, msg.Text); err != nil {
			b.log.Println("OnText(send message)", err)
		}
		return nil
	}
}
