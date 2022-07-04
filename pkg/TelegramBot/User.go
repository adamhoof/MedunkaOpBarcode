package telegrambot

import "strconv"

type User struct {
	Id string
}

func (u *User) Recipient() string {
	return u.Id
}

func (u *User) IDAsInt() int64 {
	id, err := strconv.Atoi(u.Id)
	if err != nil {
		return 0
	}
	return int64(id)
}
