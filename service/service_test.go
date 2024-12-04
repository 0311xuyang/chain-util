package service_test

import (
	"testing"

	"github.com/0311xuyang/chain-util/service"
)

func TestService(t *testing.T) {
	svc := service.Service{}
	svc.DeleteCoin("ATM")
	err := svc.CreateCoin("ATM", "1.414")
	if err != nil {
		t.Error("Error creating coin:", err)
	}
	coin, err := svc.GetCoin("ATM")
	if err != nil {
		t.Error("Error getting coin:", err)
	}
	if coin.Symbol != "ATM" {
		t.Error("Coin symbol does not match.")
	}
}
