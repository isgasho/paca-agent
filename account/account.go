package account

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
)

type Client struct {
	AlpacaClient *alpaca.Client
	Account      *alpaca.Account
}

func InitializeClient() (*Client, error) {
	// paper-trading
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
	// prod
	// alpaca.SetBaseUrl("https://api.alpaca.markets")

	// Set credentials as environment variables
	// APCA_API_KEY_ID=
	// APCA_API_SECRET_KEY=
	alpacaClient := alpaca.NewClient(common.Credentials())

	acct, err := alpacaClient.GetAccount()
	if err != nil {
		return nil, err
	}

	Client := &Client{AlpacaClient: alpacaClient, Account: acct}

	return Client, nil
}

func (c *Client) GetAccount() *alpaca.Account {
	acct := c.Account

	return acct
}
