package app

import (
	"time"

	"github.com/b2r2/tg-admin-changer/internal/models"

	"github.com/sirupsen/logrus"
	tele "gopkg.in/tucnak/telebot.v3"
)

type (
	bot struct {
		log            *logrus.Logger
		bot            *tele.Bot
		admins         map[int64]struct{}
		inlineMainMenu *tele.ReplyMarkup
		inlinePrevMenu *tele.ReplyMarkup
	}
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
	admins[237426682] = struct{}{}  // me
	admins[1087968824] = struct{}{} // group ID
	admins[666581102] = struct{}{}

	mm := &tele.ReplyMarkup{}
	inlineBtnPrice := mm.Data(models.OnBtnPrice, models.OnBtnPrice)
	inlineContacts := mm.Data(models.OnContacts, models.OnContacts)

	mm.Inline(mm.Row(inlineBtnPrice), mm.Row(inlineContacts))

	pm := &tele.ReplyMarkup{}
	inlinePrev := pm.Data(models.OnPrev, models.OnPrev)
	pm.Inline(pm.Row(inlinePrev))

	return &bot{bot: b, log: log, admins: admins, inlineMainMenu: mm, inlinePrevMenu: pm}, nil
}
