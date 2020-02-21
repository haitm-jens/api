package domain

type User struct {
	ID                uint
	Email             string
	EncryptedPassword string
	ConfirmedUnixtime int
}

func (*User) TableName() string {
	return "users"
}
