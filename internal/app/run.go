package app

import (
	tele "gopkg.in/tucnak/telebot.v3"
)

func (b *bot) Run() {
	b.log.Printf("start bot[%s]\n", b.bot.Token)
	var (
		m              = &tele.ReplyMarkup{}
		inlineBtnPrice = m.Data(onBtnPrice, onBtnPrice)
		inlineContacts = m.Data(onContacts, onContacts)
	)

	m.Inline(m.Row(inlineBtnPrice), m.Row(inlineContacts))
	menu = m

	b.bot.Handle(onStart, b.OnStart())
	b.bot.Handle(tele.OnText, b.OnText())
	b.bot.Handle(tele.OnCallback, b.OnCallback())

	go b.bot.Start()
}
