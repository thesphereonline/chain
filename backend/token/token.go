package token

import (
	"math/big"
)

type Token struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	TotalSupply *big.Int `json:"totalSupply"`
	Balance  *big.Int `json:"balance"`
}

func NewToken() *Token {
	return &Token{
		Name:     "SPHERE Coin",
		Symbol:   "SPHERE",
		TotalSupply: big.NewInt(1000000),
		Balance:  big.NewInt(0),
	}
}

func (t *Token) Transfer(amount *big.Int) {
	t.Balance.Add(t.Balance, amount)
}

func (t *Token) GetBalance() *big.Int {
	return t.Balance
}