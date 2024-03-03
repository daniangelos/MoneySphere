package models

import (
	"fmt"

	"github.com/MoneySphere/constants"
)

type Client struct {
	ID           int
	AccountLimit int
	Balance      int
}

func (c *Client) UpdateBalance(transactionType constants.TypeEnum, value int) error {
	switch transactionType {
	case constants.CreditType:
		c.creditTransaction(value)
		return nil
	case constants.DebitType:
		return c.debitTransaction(value)
	default:
		return fmt.Errorf("not valid transaction type")
	}
}

func (c *Client) creditTransaction(value int) {
	c.Balance += value
}

func (c *Client) debitTransaction(value int) error {
	if c.Balance-value < -c.AccountLimit {
		return constants.ErrNotEnoughBalance
	}
	c.Balance -= value
	return nil
}
