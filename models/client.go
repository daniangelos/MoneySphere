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
		return c.creditTransaction(value)
	case constants.DebitType:
		return c.debitTransaction(value)
	default:
		return fmt.Errorf("not valid transaction type")
	}
}

func (c *Client) creditTransaction(value int) error {
	return nil
}

func (c *Client) debitTransaction(value int) error {
	return nil
}
