package mongoInterface

import (
	"context"
)

type GenericMongoTemplate[T any] interface {
	findById(context.Context, string) (T, error)
}
