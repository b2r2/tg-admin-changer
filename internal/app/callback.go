package app

import (
	"fmt"
	"strings"

	"github.com/b2r2/tg-admin-changer/internal/config"
	"github.com/b2r2/tg-admin-changer/internal/models"

	tele "gopkg.in/tucnak/telebot.v3"
)

func (b *bot) OnCallback() tele.HandlerFunc {
	return func(c tele.Context) error {
		cb := c.Callback()
		message := ""

		if err := b.bot.Respond(cb, &tele.CallbackResponse{
			CallbackID: cb.ID,
			Text:       "Сообщение обновлено",
			ShowAlert:  false,
			URL:        "",
		}); err != nil {
			b.log.Println(err)
		}

		switch {
		case strings.Contains(cb.Data, onBtnPrice):
			message = config.Get().GetMapping()[models.Pricing].String()
			message += "\n\nПродолжайте писать в чат. Мы с Вами свяжемся."
		case strings.Contains(cb.Data, onContacts):
			message = config.Get().GetMapping()[models.Contacts].String()
		case strings.Contains(cb.Data, onPrev):
			message = config.Get().GetMapping()[models.Greeting].String()
		}

		m := &tele.ReplyMarkup{}
		inlinePrev := m.Data(onPrev, onPrev)
		m.Inline(m.Row(inlinePrev))
		_, err := b.bot.Edit(c.Message(), fmt.Sprintf(c.Message().Text+"\n\n"+message), &tele.SendOptions{
			DisableWebPagePreview: true,
			ParseMode:             tele.ModeMarkdown,
		})
		if err != nil {
			b.log.Println("OnCallback", err)
		}

		return nil
	}
}
