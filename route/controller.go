package route

import (
	"net/http"

	"github.com/0311xuyang/chain-util/db"
	"github.com/0311xuyang/chain-util/service"
	"github.com/gin-gonic/gin"
)

// GetUser gets a user.
func GetUser(c *gin.Context) {
	var request GetUserRequest
	err := c.BindUri(request)
	if err != nil {
		formatBadResponse(c, err.Error())
	}
	store := db.New()
	user, err := store.GetUser(request.ID)
	if err != nil {
		formatBadResponse(c, err.Error())
	}
	formatOKResponse(c, "success", user)
}

// SetUser sets a user.
func SetUser(c *gin.Context) {
	var request SetUserRequest
	c.BindJSON(request)
	store := db.New()
	err := store.PutUser(&db.User{
		ID:    request.ID,
		Name:  request.Name,
		Email: request.Email,
	})
	if err != nil {
		formatBadResponse(c, err.Error())
	}
	formatOKResponse(c, "success", nil)
}

// GetCoinPrice gets the price of a coin.
func GetCoinPrice(c *gin.Context) {
	var request GetCoinPriceRequest
	err := c.BindUri(request)
	if err != nil {
		formatBadResponse(c, err.Error())
	}
	svc := service.Service{}
	coin, err := svc.GetCoin(request.Symbol)
	if err != nil {
		formatBadResponse(c, err.Error())
	}
	formatOKResponse(c, "success", coin)
}

// SetCoinPrice sets the price of a coin.
func SetCoinPrice(c *gin.Context) {
	var request SetCoinPriceRequest
	c.BindJSON(request)
	svc := service.Service{}
	err := svc.CreateCoin(request.Symbol, request.Price)
	if err != nil {
		formatBadResponse(c, err.Error())
	}
	formatOKResponse(c, "success", nil)
}

func formatOKResponse(c *gin.Context, message string, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}

func formatBadResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
}
