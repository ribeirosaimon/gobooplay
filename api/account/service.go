package account

import (
	"context"
	"errors"
	"log"
	"ribeirosaimon/gobooplay/api/repository"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/security"
	"time"
)

type accountService struct {
	repository repository.Account
}

func service() accountService {
	return accountService{repository: repository.NewAccountRepository()}
}

func (s accountService) saveAccountService(ctx context.Context, dto domain.AccountDTO) (domain.AccountDTO, error) {
	existUser := s.repository.ExistUserWithLogin(ctx, dto.Login)
	if existUser {
		return domain.AccountDTO{}, errors.New("account already exists")
	}

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
	newAccount.Role = append(newAccount.Role, domain.ADMIN)
	newAccount.CreatedAt = time.Now()

	newAccount.Password = string(encriptedPassword)

	_, err = s.repository.SaveAccount(context.Background(), &newAccount)
	if err != nil {
		return domain.AccountDTO{}, err
	}
	return dto, nil
}

func (s accountService) login(ctx context.Context, login domain.LoginDTO) (domain.UserAccessToken, error) {
	currentTime := time.Now()
	var acessToken = domain.UserAccessToken{}

	account, err := s.repository.FindAccountByLogin(ctx, login.Login)
	if err != nil {
		return acessToken, err
	}

	if account.PasswordErrorCount >= 10 {
		then := account.LastLoginAttemp.Add(time.Hour * 1)
		before := currentTime.After(then)
		if !before {
			return acessToken, errors.New("you exceeded your attempts")
		}
	}

	if err := security.VerifyPassword(account.Password, login.Password); err != nil {
		account.LastLoginAttemp = currentTime
		account.PasswordErrorCount += 1
		_, err := s.repository.SaveAccount(ctx, &account)
		if err != nil {
			return acessToken, err
		}
		return acessToken, err
	}

	account.PasswordErrorCount = 0
	account.LoginCount += 1
	account.LastLogin = currentTime

	s.repository.SaveAccount(ctx, &account)

	token, err := security.CreateToken(account)
	if err != nil {
		return acessToken, err
	}
	acessToken.Token = token
	return acessToken, nil
}
