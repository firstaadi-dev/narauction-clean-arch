package domain

import "time"

type Credential struct {
	ID             uint       `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	DeletedAt      *time.Time `json:"deletedAt"`
	Username       string     `json:"username"`
	HashedPassword string     `json:"hashedPassword"`
}

type Key struct {
	Token string `json:"token"`
}

type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthUsecase interface {
	Login(userCreds UserCredential) (*Key, error)
}

type AuthRepository interface {
	Login(userCreds UserCredential) (res *Key, err error)
	IsPasswordMatch(password string, hashedPassword string) (err error)
}
