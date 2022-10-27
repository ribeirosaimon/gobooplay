package main

import (
	"context"
	"fmt"
	"ribeirosaimon/gobooplay/api/repository"
	"ribeirosaimon/gobooplay/domain"
)

func main() {
	meuTeste := domain.Account{Name: "Meu teste"}
	id, err := repository.MongoTemplate[domain.Account]().Save(context.Background(), meuTeste)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)

	//r := gin.Default()
	//
	//routers.CreateConfigRouter(r)
	//
	//r.Run(":8080")
}
