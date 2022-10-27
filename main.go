package main

import (
	"context"
	"fmt"
	"ribeirosaimon/gobooplay/api/repository"
	"ribeirosaimon/gobooplay/domain"
)

func main() {
	value := "6356062d9f3ea305167062ba"
	id, err := repository.NewMongoTemplateRepository(domain.Account{}).FindById(context.Background(), value)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)

	//r := gin.Default()
	//
	//routers.CreateConfigRouter(r)
	//
	//r.Run(":8080")
}
