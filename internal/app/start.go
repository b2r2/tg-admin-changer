package app

import (
	"bytes"
	"context"
	text "text/template"

	"github.com/b2r2/tg-admin-changer/internal/config"
	"github.com/b2r2/tg-admin-changer/internal/models"
	tele "gopkg.in/tucnak/telebot.v3"
)

func (b *bot) Start(ctx context.Context) {
	// TODO
	b.log.Println("start bot[", b.bot.Token, "]")
	var (
		m        = &tele.ReplyMarkup{ResizeKeyboard: true}
		btnPrice = m.Text("💲 Цены")
	)

	m.Reply(m.Row(btnPrice))

	menu = m

	b.bot.Handle("/start", func(c tele.Context) error {
		u, err := b.GetUser(ctx, c.Chat().Username)
		if err != nil {
			b.log.Println("error get user:", err)
		}
		var msg = new(bytes.Buffer)
		if u == nil {
			msg = config.Get().GetMapping()[models.Greeting]
			err = b.SetUser(ctx, c.Chat().FirstName, c.Chat().Username)
			if err != nil {
				b.log.Println("error set user:", err)
			}
		} else {
			if err = b.UpdateUser(ctx, u.ID); err != nil {
				b.log.Println("error update user:", err)
			}

			// TODO add new layer parse message
			t, err := text.New("text").Parse(config.Get().GetMapping()[models.Repeated].String())
			if err != nil {
				return err
			}
			data := map[string]interface{}{
				"first":      u.First,
				"updated_at": u.UpdatedAt.Format("2006/01/02 15:04:05"),
			}
			if err := t.Execute(msg, data); err != nil {
				return err
			}
		}
		return b.SendMessage(c.Chat().ID, msg.String())
	})

	b.bot.Handle(&btnPrice, func(c tele.Context) error {
		price := config.Get().GetMapping()[models.Pricing]
		return b.SendMessage(c.Chat().ID, price.String())
	})

	// TODO
	// reply message on common admin chat
	b.bot.Handle(tele.OnText, func(c tele.Context) error {
		_, err := b.bot.Forward(&tele.Chat{ID: b.admins[0].chadId}, c.Message())
		if err != nil {
			b.log.Println(err)
		}

		return nil
	})

	go b.bot.Start()

	<-ctx.Done()
}
