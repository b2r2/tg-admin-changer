package app

import (
	"github.com/b2r2/tg-admin-changer/internal/models"
	tele "gopkg.in/tucnak/telebot.v3"
)

const (
	negativeMessage = `❗Сообщение не отправлено❗ 
Необходимо выделить сообщение кому хочешь написать`
	negativeMessageToBot = `❗Сообщение не отправлено❗ 
Невозможно отправить сообщение боту`
)

func (b *bot) OnText() tele.HandlerFunc {
	return func(c tele.Context) error {
		m := &tele.ReplyMarkup{}
		inlinePrev := m.Data(models.OnContacts, models.OnContacts)
		m.Inline(m.Row(inlinePrev))

		cid := c.Sender().ID

		if _, ok := b.admins[cid]; !ok {
			_, err := b.bot.Forward(&tele.Chat{ID: models.Channel}, c.Message(), &tele.SendOptions{
				ParseMode: tele.ModeMarkdown,
			})
			if err != nil {
				b.log.Println("OnText(forward message)", err)
			}
			return nil
		}

		if c.Message().ReplyTo == nil {
			_, err := b.bot.Send(&tele.Chat{ID: models.Channel}, negativeMessage)
			if err != nil {
				b.log.Println("OnText(forward message)", err)
			}
			return nil
		}

		if c.Message().ReplyTo.OriginalSender == nil {
			_, err := b.bot.Send(&tele.Chat{ID: models.Channel}, negativeMessageToBot)
			if err != nil {
				b.log.Println("OnText(forward message)", err)
			}
			return nil
		}

		id := c.Message().ReplyTo.OriginalSender.ID
		_, err := b.bot.Send(&tele.Chat{ID: id}, c.Message().Text, &tele.SendOptions{ReplyMarkup: m})
		if err != nil {
			b.log.Println("OnText(send message)", err)
		}
		return nil
	}
}
