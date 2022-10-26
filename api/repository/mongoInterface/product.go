package mongoInterface

import (
	"context"
	"ribeirosaimon/gobooplay/domain"
)

type Product interface {
	Save(context.Context, *domain.Product) (domain.Product, error)
	FindById(context.Context, string) (domain.Product, error)
	DelebeById(context.Context, string) error
}
