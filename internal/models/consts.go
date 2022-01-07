package models

import "bytes"

const (
	Greeting = "greeting"
	Pricing  = "price"
	Contacts = "contacts"
)

var Texts = map[string]*bytes.Buffer{
	Greeting: bytes.NewBufferString(greeting),
	Pricing:  bytes.NewBufferString(price),
	Contacts: bytes.NewBufferString(contact),
}
