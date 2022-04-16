package usecase

import "github.com/firstaadi-dev/narauction-clean-arch/domain"

type AuthUsecase struct {
	AuthRepository domain.AuthRepository
}

// Login implements domain.AuthUsecase
func (a *AuthUsecase) Login(userCreds domain.UserCredential) (*domain.Key, error) {
	res, err := a.AuthRepository.Login(userCreds)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func NewAuthUsecase(a domain.AuthRepository) domain.AuthUsecase {
	return &AuthUsecase{
		AuthRepository: a,
	}
}
