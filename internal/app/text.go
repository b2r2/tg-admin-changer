package app

import (
	tele "gopkg.in/tucnak/telebot.v3"
)

func (b *bot) OnText() tele.HandlerFunc {
	return func(c tele.Context) error {
		m := &tele.ReplyMarkup{}
		inlinePrev := m.Data(onContacts, onContacts)
		m.Inline(m.Row(inlinePrev))

		cid := c.Message().Sender.ID
		if b.admins[0].chadId != cid && b.admins[1].chadId != cid {
			_, err := b.bot.Forward(&tele.Chat{ID: channel}, c.Message(), &tele.SendOptions{
				ParseMode: tele.ModeMarkdown,
			})
			if err != nil {
				b.log.Println("OnText(forward message)", err)
			}
			return nil
		}

		if c.Update().Message.ReplyTo.OriginalSender == nil {
			return nil
		}
		id := c.Update().Message.ReplyTo.OriginalSender.ID
		_, err := b.bot.Send(&tele.Chat{ID: id}, c.Message().Text, &tele.SendOptions{ReplyMarkup: m})
		if err != nil {
			b.log.Println("OnText(send message)", err)
		}
		return nil
	}
}
