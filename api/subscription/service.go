package subscription

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
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
	rest := subscription.EndAt.Sub(now)
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

	now := time.Now()
	addDate := now.Add(time.Duration(subscription.RestOfSubscription))
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

func (s SubscriptionService) getRestOfSubscription(c *gin.Context, user domain.LoggedUser) (error, string) {
	subscription, err := s.FindSubscription(c, user)
	if err != nil {
		return err, ""
	}

	if subscription.Status == domain.PAUSE {
		rest := time.Duration(subscription.RestOfSubscription)
		return nil, fmt.Sprintf("You have: %f Days and %f Hours", rest.Hours()/24, rest.Hours())
	}
	until := time.Until(subscription.EndAt)
	var days string
	if until.Hours() <= 1 {
		days = fmt.Sprintf("You have: %f Days and ", until.Hours()/24)
	}
	return nil, fmt.Sprintf("%s %f Hours", days, until.Hours())
}
