package consumerApi

import (
	"ribeirosaimon/gobooplay/domain"
	"ribeirosaimon/gobooplay/util"
)

func GetApiMovieInformations() []domain.Movies {
	movies := []domain.Movies{
		{Name: "GOT", Hash: util.CreateHash()},
		{Name: "Ring of Power", Hash: util.CreateHash()},
	}

	return movies
}
