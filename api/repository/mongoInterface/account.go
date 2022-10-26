package mongoInterface

import (
	"context"
	"ribeirosaimon/gobooplay/domain"
)

type Account interface {
	Save(context.Context, *domain.Account) (domain.Account, error)
	FindById(context.Context, string) (domain.Account, error)
	DelebeById(context.Context, string) error
	ExistUserWithLogin(context.Context, string) bool
	FindAccountByLogin(context.Context, string) (domain.Account, error)
}
