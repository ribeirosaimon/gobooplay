package movies

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ribeirosaimon/gobooplay/exceptions"
	"ribeirosaimon/gobooplay/util"
)

type controllerMovie struct {
	service movieService
}

func ControllerMovie() controllerMovie {
	return controllerMovie{
		service: ServiceMovie(),
	}
}

func (s controllerMovie) GetMovie(c *gin.Context) {
	user := util.GetUser(c)
	movies, err := s.service.getAllMovies(c, user)

	if err != nil {
		exceptions.ValidateException(c, err.Error(), http.StatusConflict)
		return
	}

	c.JSON(http.StatusOK, movies)
	return
}
