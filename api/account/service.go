package account

import (
	"ribeirosaimon/gobooplay/api/repository"
	"ribeirosaimon/gobooplay/domain"
)

type accountService struct {
	repository repository.Account
}

func service() accountService {
	return accountService{repository: repository.NewAccountRepository()}
}

func (s accountService) saveAccount() {
	s.repository.SaveAccount(domain.Account{})
}
