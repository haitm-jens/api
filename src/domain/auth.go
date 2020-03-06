package domain

type Auth struct {
	ID                uint
	Email             string
	EncryptedPassword string
	ConfirmedUnixtime int
}

func (Auth) TableName() string {
	return "users"
}
