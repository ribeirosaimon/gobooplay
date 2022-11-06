package subscription

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/repository"
	"time"
)

type SubscriptionService struct {
	subscriptionRepository repository.MongoTemplateStruct[domain.Subscription]
}

func ServiceSubscription() SubscriptionService {
	return SubscriptionService{
		subscriptionRepository: repository.MongoTemplate[domain.Subscription](),
	}
}

func (s SubscriptionService) CreateSubscription(c context.Context, user domain.Account, product domain.Product) (
	domain.Subscription, error) {
	var mySubs domain.Subscription
	now := time.Now()
	mySubs.Product = product
	mySubs.CreatedAt = time.Now()
	mySubs.UpdatedAt = time.Now()
	mySubs.Owner = user.MyRef()
	mySubs.BegginAt = now
	mySubs.EndAt = now.AddDate(0, int(product.SubscriptionTime), 0)
	mySubs.Status = domain.ACTIVE
	mySubs, err := s.subscriptionRepository.Save(c, mySubs)
	if err != nil {
		return domain.Subscription{}, err
	}
	return mySubs, nil
}

func (s SubscriptionService) FindSubscription(c context.Context, user domain.LoggedUser) (domain.Subscription, error) {
	filter := bson.D{
		{"owner.userId", user.UserId},
	}
	subs, err := s.subscriptionRepository.FindOneByFilter(c, filter)
	if err != nil {
		return domain.Subscription{}, err
	}
	return subs, nil
}

func (s SubscriptionService) PauseSubscription(c context.Context, user domain.LoggedUser) error {
	subscription, err := s.FindSubscription(c, user)
	if err != nil {
		return err
	}
	now := time.Now()
	subscription.PauseAt = now
	rest := now.Sub(subscription.EndAt)
	filter := bson.D{
		{"restOfSubscription", rest},
		{"pauseAt", now},
		{"status", domain.PAUSE},
	}

	_, err = s.subscriptionRepository.UpdateById(c, subscription.ID.Hex(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (s SubscriptionService) ValidateSubscription(c context.Context, user domain.LoggedUser) error {
	subscription, err := s.FindSubscription(c, user)
	if err != nil {
		return err
	}
	now := time.Now()
	if now.After(subscription.EndAt) {
		return errors.New("your subscribe is expired")
	}
	return nil
}

func (s SubscriptionService) ActivateSubscription(c context.Context, user domain.LoggedUser) error {
	subscription, err := s.FindSubscription(c, user)
	if err != nil {
		return err
	}

	addDate := subscription.EndAt.Add(time.Duration(subscription.RestOfSubscription))
	filter := bson.D{
		{"endAt", addDate},
		{"status", domain.ACTIVE},
	}

	_, err = s.subscriptionRepository.UpdateById(c, subscription.ID.Hex(), filter)
	if err != nil {
		return err
	}
	return nil
}
