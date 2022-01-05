package app

import (
	"time"

	"github.com/sirupsen/logrus"
	tele "gopkg.in/tucnak/telebot.v3"
)

const (
	onBtnPrice = "Наши каналы и цены"
	onContacts = "Наши контакты"
	onPrev     = "Назад"
	onStart    = "/start"
	channel    = -1001383844955
)

type (
	bot struct {
		log    *logrus.Logger
		bot    *tele.Bot
		admins admins
	}
	admin struct {
		chadId int64
	}
	admins []admin
)

var (
	menu *tele.ReplyMarkup
)

func New(log *logrus.Logger, t string) (*bot, error) {
	b, err := tele.NewBot(tele.Settings{
		Token:     t,
		Poller:    &tele.LongPoller{Timeout: time.Second * 10},
		ParseMode: tele.ModeMarkdown,
	})

	if err != nil {
		return nil, err
	}

	as := admins{
		{chadId: 237426682}, // me
		{chadId: 666581102},
	}

	return &bot{bot: b, log: log, admins: as}, nil
}
