package models

import "bytes"

const (
	Greeting   = "greeting"
	Pricing    = "price"
	Contacts   = "contacts"
	OnBtnPrice = "Наши каналы и цены"
	OnContacts = "Наши контакты"
	OnPrev     = "Назад"
	OnStart    = "/start"
	Filename   = ".env"
	Channel    = -1001707672035
	//Channel  = -1001383844955 // debug
	//Filename = "TOKEN"

	NegativeMessage = `❗Сообщение не отправлено❗ 
Необходимо выделить сообщение кому хочешь написать`
	NegativeMessageToBot = `❗Сообщение не отправлено❗ 
Невозможно отправить сообщение боту`
	SelfMessage = `❗Сообщение не отправлено❗
Вероятно потому что Вы администратор, если нет, то лучше напишите напрямую администраторам каналов`
)

var Texts = map[string]*bytes.Buffer{
	Greeting: bytes.NewBufferString(greeting),
	Pricing:  bytes.NewBufferString(price),
	Contacts: bytes.NewBufferString(contact),
}
