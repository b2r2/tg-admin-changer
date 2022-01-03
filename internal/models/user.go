package models

import (
	"time"

	jsoniter "github.com/json-iterator/go"
)

type (
	User struct {
		ID        uint64     `db:"id" json:"id"`
		First     string     `db:"first" json:"first"`
		Username  string     `db:"username" json:"username"`
		CreatedAt *time.Time `db:"created_at" json:"created_at"`
		UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	}
	Users []User
)

func (u User) MarshalBinary() ([]byte, error) {
	return jsoniter.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return jsoniter.Unmarshal(data, &u)
}
