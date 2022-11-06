package movies

import (
	"context"
	"ribeirosaimon/gobooplay/api/subscription"
	"ribeirosaimon/gobooplay/consumerApi"
	"ribeirosaimon/gobooplay/domain"
)

type movieService struct {
	allMovies           []domain.Movies
	subscriptionService subscription.SubscriptionService
}

func ServiceMovie() movieService {
	return movieService{
		allMovies:           consumerApi.GetApiMovieInformations(),
		subscriptionService: subscription.ServiceSubscription(),
	}
}

func (s movieService) getAllMovies(c context.Context, user domain.LoggedUser) ([]domain.Movies, error) {
	if err := s.subscriptionService.ValidateSubscription(c, user); err != nil {
		return []domain.Movies{}, err
	}

	return s.allMovies, nil
}
