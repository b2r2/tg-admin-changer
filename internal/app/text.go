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

		message := ""
		opts := &tele.SendOptions{}
		switch {
		case msg.ReplyTo == nil:
			if _, ok := b.admins[msg.Chat.ID]; ok {
				cid = msg.Sender.ID
				message = models.SelfMessage
				opts.ReplyMarkup = b.inlineMainMenu
			}
			if msg.Chat.ID == models.Channel {
				cid = models.Channel
				message = models.NegativeMessage
			}
		case msg.ReplyTo != nil && msg.ReplyTo.OriginalSender == nil:
			if msg.ReplyTo.Chat.ID == models.Channel {
				cid = models.Channel
				message = models.NegativeMessageToBot
			}
		default:
			cid = msg.ReplyTo.OriginalSender.ID
			message = msg.Text
		}

		if _, err := b.bot.Send(&tele.Chat{ID: cid}, message, opts); err != nil {
			b.log.Println("OnText(send message)", err)
		}
		return nil
	}
}
