package app

import (
	"time"

	"github.com/b2r2/tg-admin-changer/internal/repositories"

	"github.com/sirupsen/logrus"
	tele "gopkg.in/tucnak/telebot.v3"
)

type (
	bot struct {
		log    *logrus.Logger
		bot    *tele.Bot
		repo   repositories.Repository
		admins []admin
	}
	admin struct {
		username string
		chadId   int64
	}
)

var menu *tele.ReplyMarkup

func New(log *logrus.Logger, t string, r repositories.Repository) (*bot, error) {
	b, err := tele.NewBot(tele.Settings{
		Token:     t,
		Poller:    &tele.LongPoller{Timeout: time.Second * 8},
		ParseMode: tele.ModeMarkdown,
	})

	if err != nil {
		return nil, err
	}

	admins := []admin{
		{
			chadId:   237426682,
			username: "ramil",
		},
		{
			// TODO replace
			chadId:   237426682,
			username: "ramil",
		},
	}

	return &bot{bot: b, repo: r, log: log, admins: admins}, nil
}
