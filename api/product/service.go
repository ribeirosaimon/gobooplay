package product

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ribeirosaimon/gobooplay/api/repository"
	"ribeirosaimon/gobooplay/domain"
	"time"
)

type productService struct {
	productRepository repository.MongoTemplateStruct[domain.Product]
}

func ServiceProduct() productService {
	return productService{
		productRepository: repository.MongoTemplate[domain.Product](),
	}
}

func (s productService) AddProduct(ctx context.Context, payload domain.ProductDTO, user domain.LoggedUser) (domain.Product, error) {
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

func (s productService) FindAllProduct(ctx context.Context, user domain.LoggedUser) ([]domain.ProductDTO, error) {
	filter := bson.D{
		{"updateBy.userId", user.UserId},
	}
	find, err := s.productRepository.Find(ctx, filter)
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

func (s productService) DeleteProductById(ctx context.Context, id string, user domain.LoggedUser) error {
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
