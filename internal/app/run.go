package app

import (
	"github.com/b2r2/tg-admin-changer/internal/models"
	tele "gopkg.in/tucnak/telebot.v3"
)

func (b *bot) Run() {
	b.log.Printf("start bot[%s]", b.bot.Token)

	b.bot.Handle(models.OnStart, b.OnStart())
	b.bot.Handle(tele.OnText, b.OnText())
	b.bot.Handle(tele.OnCallback, b.OnCallback())

	go b.bot.Start()
}
