package product

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository"
	"ribeirosaimon/gobooplay/util"
	"time"
)

type ProductService struct {
	productRepository repository.MongoTemplateStruct[domain.Product]
}

func ServiceProduct() ProductService {
	return ProductService{
		productRepository: repository.MongoTemplate[domain.Product](),
	}
}

func (s ProductService) GetTrialProduct(c context.Context) (domain.Product, error) {
	filter := bson.D{
		{"status", domain.TRIAL},
	}
	product, err := s.productRepository.FindOneByFilter(c, filter)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s ProductService) AddProduct(ctx context.Context, payload domain.ProductDTO, user domain.LoggedUser) (domain.Product, error) {
	price, err := primitive.ParseDecimal128(payload.Price)
	if err != nil {
		return domain.Product{}, err
	}

	var product domain.Product

	product.Name = payload.Name
	product.Price = price
	product.SubscriptionTime = payload.SubscriptionTime
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	product.UpdateBy = user

	savedProduct, err := s.productRepository.Save(ctx, product)
	if err != nil {
		return domain.Product{}, err
	}
	return savedProduct, nil
}

func (s ProductService) FindAllProduct(ctx context.Context, user domain.LoggedUser) ([]domain.ProductDTO, error) {
	var filter bson.D
	if util.ContainsRole[domain.Role](user.Role, domain.ADMIN) {
		filter = bson.D{
			{"updateBy.userId", user.UserId},
		}
	} else {
		filter = bson.D{{}}
	}

	find, err := s.productRepository.FindAllByFilter(ctx, filter)
	if err != nil {
		return []domain.ProductDTO{}, nil
	}
	var sliceOfProduct []domain.ProductDTO

	for _, product := range find {
		var productConverter domain.ProductDTO
		productConverter.ID = product.ID.Hex()
		productConverter.Name = product.Name
		productConverter.SubscriptionTime = product.SubscriptionTime
		productConverter.Price = product.Price.String()

		sliceOfProduct = append(sliceOfProduct, productConverter)

	}
	return sliceOfProduct, nil
}

func (s ProductService) DeleteProductById(ctx context.Context, id string, user domain.LoggedUser) error {
	product, err := s.productRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if product.UpdateBy.UserId != user.UserId {
		return errors.New("you not permission")
	}

	if err := s.productRepository.DeleteById(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s ProductService) UpdateProduct(c *gin.Context, payload domain.ProductDTO, id string, user domain.LoggedUser) (domain.Product, error) {
	product, err := s.productRepository.FindById(c, id)
	if err != nil {
		return domain.Product{}, err
	}
	if product.UpdateBy.UserId != user.UserId {
		return domain.Product{}, errors.New("you not permission")
	}

	var filterBson = bson.D{}
	if payload.Name != "" {
		filterBson = append(filterBson, bson.E{Key: "name", Value: payload.Name})
	}

	if payload.Price != "" {
		priceDecimal, err := primitive.ParseDecimal128(payload.Price)

		if err != nil {
			return domain.Product{}, err
		}
		filterBson = append(filterBson, bson.E{Key: "price", Value: priceDecimal})
	}

	if payload.SubscriptionTime != 0 {
		filterBson = append(filterBson, bson.E{Key: "subscriptionTime", Value: payload.SubscriptionTime})
	}
	filterBson = append(filterBson, bson.E{Key: "updatedAt", Value: time.Now()})

	response, err := s.productRepository.UpdateById(c, id, filterBson)
	if err != nil {
		return domain.Product{}, err
	}
	return response, nil
}
