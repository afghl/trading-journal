package main

import (
	"trading_journal/domain/service"
	"trading_journal/infrastructure/mysql"
)

func main() {
	if err := initDomainProvider(); err != nil {
		panic("provider init fail")
	}

}

func initDomainProvider() error {
	return service.InitProvider(mysql.GetRepo())
}
