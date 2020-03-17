package domain

type User struct {
	ID                uint
	Email             string
	ConfirmedUnixtime uint
}

func (*User) TableName() string {
	return "users"
}
