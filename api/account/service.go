package account

import (
	"context"
	"ribeirosaimon/gobooplay/api/repository"
	"ribeirosaimon/gobooplay/domain"
)

type accountService struct {
	repository repository.Account
}

func service() accountService {
	return accountService{repository: repository.NewAccountRepository()}
}

func (s accountService) saveAccount(dto domain.AccountDTO) (domain.Account, error) {
	var newAccount domain.Account

	newAccount.Name = dto.Name
	newAccount.FamilyName = dto.FamilyName
	newAccount.Gender = dto.Gender
	newAccount.Login = dto.Login
	newAccount.Password = dto.Password

	account, err := s.repository.SaveAccount(context.Background(), &newAccount)
	if err != nil {
		return domain.Account{}, err
	}
	return account, nil
}
