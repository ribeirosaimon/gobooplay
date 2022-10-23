package account

import (
	"context"
	"log"
	"ribeirosaimon/gobooplay/api/repository"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/security"
)

type accountService struct {
	repository repository.Account
}

func service() accountService {
	return accountService{repository: repository.NewAccountRepository()}
}

func (s accountService) saveAccount(ctx context.Context, dto domain.AccountDTO) (domain.AccountDTO, error) {
	var newAccount domain.Account
	encriptedPassword, err := security.EncriptyPassword(dto.Password)
	if err != nil {
		log.Printf(err.Error())
		return domain.AccountDTO{}, err
	}

	newAccount.Name = dto.Name
	newAccount.FamilyName = dto.FamilyName
	newAccount.Gender = dto.Gender
	newAccount.Login = dto.Login

	newAccount.Password = string(encriptedPassword)

	_, err = s.repository.SaveAccount(context.Background(), &newAccount)
	if err != nil {
		return domain.AccountDTO{}, err
	}
	return dto, nil
}

func (s accountService) login(ctx context.Context, login domain.LoginDTO) (domain.UserAccessToken, error) {
	var acessToken = domain.UserAccessToken{}

	account, err := s.repository.FindAccountByLogin(ctx, login.Login)
	if err != nil {
		return acessToken, err
	}

	if err := security.VerifyPassword(account.Password, login.Password); err != nil {
		return acessToken, err
	}

	token, err := security.CreateToken(account)
	if err != nil {
		return acessToken, err
	}
	acessToken.Token = token
	return acessToken, nil
}
