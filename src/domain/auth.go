package domain

type Auth struct {
	ID                uint
	Email             string
	EncryptedPassword string
	ConfirmedUnixtime uint
}

func (Auth) TableName() string {
	return "users"
}
