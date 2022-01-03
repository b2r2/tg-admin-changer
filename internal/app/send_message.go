package app

import (
	tele "gopkg.in/tucnak/telebot.v3"
)

func (b *bot) SendMessage(cid int64, to interface{}) error {
	_, err := b.bot.Send(&tele.Chat{ID: cid}, to, &tele.SendOptions{
		ReplyMarkup:           menu,
		DisableWebPagePreview: true,
		ParseMode:             tele.ModeMarkdown,
	})

	return err
}
