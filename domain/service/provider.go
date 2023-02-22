package service

import (
	"errors"
	"trading_journal/domain/repository"
)

// provider 封装领域服务所需要使用的外部依赖的封装
type provider struct {
	Repo repository.TradeRepository
}

func getProvider() *provider {
	return globalProvider
}

var globalProvider *provider

func InitProvider(repo repository.TradeRepository) error {
	if repo == nil {
		return errors.New("not provided")
	}
	globalProvider = &provider{
		Repo: repo,
	}
	return nil
}
