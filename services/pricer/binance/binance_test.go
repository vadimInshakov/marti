package binance

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vadimInshakov/marti/entity"
	"os"
	"testing"
)

func Test(t *testing.T) {
	apikey := os.Getenv("APIKEY")
	require.NotEmpty(t, apikey, "APIKEY env is not set")
	secretkey := os.Getenv("SECRETKEY")
	require.NotEmpty(t, apikey, "SECRETKEY env is not set")
	pricer := NewPricer(apikey, secretkey)
	price, err := pricer.GetPrice(entity.Pair{From: "BTC", To: "RUB"})
	assert.NoError(t, err)
	f, _ := price.Float64()
	assert.Greater(t, f, 1.0)
}
