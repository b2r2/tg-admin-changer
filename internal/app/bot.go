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
	channel    = -1001707672035
	//channel = -1001383844955 // debug
)

type (
	bot struct {
		log    *logrus.Logger
		bot    *tele.Bot
		admins map[int64]struct{}
	}
)

var (
	menu *tele.ReplyMarkup
)

func New(log *logrus.Logger, token string) (*bot, error) {
	b, err := tele.NewBot(tele.Settings{
		Token:     token,
		Poller:    &tele.LongPoller{Timeout: time.Second * 10},
		ParseMode: tele.ModeMarkdown,
	})

	if err != nil {
		return nil, err
	}

	admins := make(map[int64]struct{})
	admins[237426682] = struct{}{} // me
	admins[666581102] = struct{}{}

	return &bot{bot: b, log: log, admins: admins}, nil
}
