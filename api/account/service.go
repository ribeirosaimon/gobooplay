package account

import (
	"context"
	"errors"
	"fmt"
	"log"
	"ribeirosaimon/gobooplay/api/order"
	"ribeirosaimon/gobooplay/api/product"
	"ribeirosaimon/gobooplay/api/subscription"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository"
	"ribeirosaimon/gobooplay/repository/mongoInterface"
	"ribeirosaimon/gobooplay/security"
	"time"
)

type accountService struct {
	accountRepository mongoInterface.Account
	orderService      order.OrderService
	productService    product.ProductService
	subscribeService  subscription.SubscriptionService
}

func service() accountService {
	return accountService{
		accountRepository: repository.NewAccountRepository(),
		orderService:      order.ServiceOrder(),
		productService:    product.ServiceProduct(),
		subscribeService:  subscription.ServiceSubscription(),
	}
}

func (s accountService) saveAccount(ctx context.Context, dto domain.AccountDTO) (domain.AccountDTO, error) {
	existUser := s.accountRepository.ExistUserWithLogin(ctx, dto.Login)
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
	newAccount.Role = append(newAccount.Role, domain.USER)
	newAccount.CreatedAt = time.Now()
	newAccount.Status = domain.ACTIVE

	newAccount.Password = string(encriptedPassword)

	account, err := s.accountRepository.Save(ctx, &newAccount)

	firstProduct, err := s.productService.GetTrialProduct(ctx)
	if err != nil {
		return domain.AccountDTO{}, err
	}
	_, err = s.subscribeService.CreateSubscription(ctx, account, firstProduct)
	if err != nil {
		return domain.AccountDTO{}, err
	}
	if err != nil {
		return domain.AccountDTO{}, err
	}
	return dto, nil
}

func (s accountService) login(ctx context.Context, login domain.LoginDTO) (domain.UserAccessToken, error) {
	currentTime := time.Now()
	var acessToken = domain.UserAccessToken{}

	account, err := s.accountRepository.FindAccountByLogin(ctx, login.Login)
	if err != nil {
		return acessToken, errors.New("this account not exist")
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
		_, err := s.accountRepository.Save(ctx, &account)
		if err != nil {
			return acessToken, errors.New("incorrect password")
		}
		return acessToken, errors.New("incorrect password")
	}

	subs, err := s.subscribeService.FindSubscription(ctx, account.GetLoggedUser())
	if err != nil {
		return domain.UserAccessToken{}, err
	}

	if subs.Status == domain.PAUSE {
		s.subscribeService.ActivateSubscription(ctx, account.GetLoggedUser())
	}

	account.PasswordErrorCount = 0
	account.LoginCount += 1
	account.LastLogin = currentTime

	s.accountRepository.Save(ctx, &account)

	token, err := security.CreateToken(account)
	if err != nil {
		return acessToken, err
	}
	acessToken.Token = token
	return acessToken, nil
}

func (s accountService) avaliableSubscription(ctx context.Context, login domain.LoginDTO) error {
	account, err := s.accountRepository.FindAccountByLogin(ctx, login.Login)
	if err != nil {
		return err
	}
	fmt.Println(account)
	return nil

}
