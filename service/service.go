package service

import (
	"time"

	"github.com/0311xuyang/chain-util/dal"
	"github.com/0311xuyang/chain-util/dal/dao"
	"github.com/0311xuyang/chain-util/dal/model"
)

// Service is a struct that holds the service methods.
type Service struct{}

// CreateCoin creates a coin.
func (s *Service) CreateCoin(symbol, price string) error {
	coin := &model.CoinPrice{
		Date:      time.Now(),
		Price:     price,
		Symbol:    symbol,
		ChainName: "HSK",
	}
	err := dao.Use(dal.DB).CoinPrice.Create(coin)
	if err != nil {
		return err
	}
	return nil
}

// GetCoin gets a coin by symbol.
func (s *Service) GetCoin(symbol string) (*model.CoinPrice, error) {
	m := dao.Use(dal.DB).CoinPrice
	coin, err := m.Select(m.ALL).Where(m.Symbol.Eq(symbol)).First()
	if err != nil {
		return nil, err
	}
	return coin, nil
}

// DeleteCoin deletes a coin by symbol.
func (s *Service) DeleteCoin(symbol string) error {
	m := dao.Use(dal.DB).CoinPrice
	_, err := m.Where(m.Symbol.Eq(symbol)).Delete()
	return err
}
