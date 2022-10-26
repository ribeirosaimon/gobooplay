package mongoInterface

import "context"

type MongoTemplate interface {
	findById(context.Context, string, interface{}) (interface{}, error)
}
