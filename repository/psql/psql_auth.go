package repository

import (
	"encoding/base64"
	"os"

	"github.com/firstaadi-dev/narauction-clean-arch/domain"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type psqlAuthRepository struct {
	DB *gorm.DB
}

// ChechPasswordHash implements domain.AuthRepository
func (p *psqlAuthRepository) IsPasswordMatch(password string, hashedPassword string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// Login implements domain.AuthRepository
func (p *psqlAuthRepository) Login(userCreds domain.UserCredential) (res *domain.Key, err error) {
	godotenv.Load()

	var Credential domain.Credential
	var Key domain.Key
	SECRET_KEY := os.Getenv("SECRET_KEY")
	p.DB.Where("username = ?", userCreds.Username).First(&Credential)
	err = p.IsPasswordMatch(userCreds.Password, Credential.HashedPassword)
	if err != nil {
		return nil, err
	}
	Key.Token = base64.StdEncoding.EncodeToString([]byte(SECRET_KEY))
	return &Key, nil
}

func NewPsqlAuthRepository(Db *gorm.DB) domain.AuthRepository {
	return &psqlAuthRepository{DB: Db}
}
