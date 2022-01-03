package models

const (
	Greeting = "text/greeting.txt"
	Pricing  = "text/price.txt"
	Repeated = "text/repeated.txt"
	KeyCache = "tg_changer:users"
)

const (
	GetUser    = "SELECT * FROM users WHERE username = $1"
	SetUser    = "INSERT INTO users(first, username) VALUES ($1, $2) RETURNING *"
	UpdateUser = "UPDATE users SET updated_at = CURRENT_TIMESTAMP WHERE id = $1"
)
