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
)

var Texts = map[string]*bytes.Buffer{
	Greeting: bytes.NewBufferString(greeting),
	Pricing:  bytes.NewBufferString(price),
	Contacts: bytes.NewBufferString(contact),
}
