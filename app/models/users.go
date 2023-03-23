package models

import "time"

type User struct {
	ID        int
	UUID      string
	NAME      string
	Email     string
	password  string
	CreatedAt time.Time
}

//ユーザの作成
func (u *User) CreateUser() (err error) {
	cmd := `insert into users(
		uuid,
		name,
		email,
		password,
		created_at) values(?,?,?,?,?)`

	_, err := Db.Exec(cmd)
}
