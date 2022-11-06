package voucher

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository"
	"time"
)

type voucherService struct {
	voucherRepository repository.MongoTemplateStruct[domain.Voucher]
}

func ServiceVoucher() voucherService {
	return voucherService{
		voucherRepository: repository.MongoTemplate[domain.Voucher](),
	}
}

func (s voucherService) AddVoucher(c context.Context, payload domain.VoucherDTO, user domain.LoggedUser) (domain.Voucher, error) {
	price, err := primitive.ParseDecimal128(payload.Price)
	if err != nil {
		return domain.Voucher{}, err
	}

	var voucher domain.Voucher

	voucher.Name = payload.Name
	voucher.Price = price
	voucher.Status = domain.ACTIVE
	voucher.Description = payload.Description
	voucher.CreatedAt = time.Now()
	voucher.UpdatedAt = time.Now()

	savedProduct, err := s.voucherRepository.Save(c, voucher)
	if err != nil {
		return domain.Voucher{}, err
	}
	return savedProduct, nil
}

func (s voucherService) deleteVoucherById(c context.Context, id string, user domain.LoggedUser) error {
	voucher, err := s.voucherRepository.FindById(c, id)
	if err != nil {
		return err
	}
	if voucher.UpdateBy.UserId != user.UserId {
		return errors.New("you not permission")
	}

	if err := s.voucherRepository.DeleteById(c, id); err != nil {
		return err
	}
	return nil
}

func (s voucherService) updateVoucher(c *gin.Context, payload domain.VoucherDTO, id string, user domain.LoggedUser) (domain.Voucher, error) {
	product, err := s.voucherRepository.FindById(c, id)
	if err != nil {
		return domain.Voucher{}, err
	}
	if product.UpdateBy.UserId != user.UserId {
		return domain.Voucher{}, errors.New("you not permission")
	}

	var filterBson = bson.D{}
	if payload.Name != "" {
		filterBson = append(filterBson, bson.E{Key: "name", Value: payload.Name})
	}

	if payload.Price != "" {
		priceDecimal, err := primitive.ParseDecimal128(payload.Price)

		if err != nil {
			return domain.Voucher{}, err
		}
		filterBson = append(filterBson, bson.E{Key: "price", Value: priceDecimal})
	}

	if payload.Description != "" {
		filterBson = append(filterBson, bson.E{Key: "description", Value: payload.Description})
	}
	filterBson = append(filterBson, bson.E{Key: "updatedAt", Value: time.Now()})

	response, err := s.voucherRepository.UpdateById(c, id, filterBson)
	if err != nil {
		return domain.Voucher{}, err
	}
	return response, nil
}
