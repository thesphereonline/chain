package api

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/thesphereonline/chain/backend/token"
)

func GetTokenBalance(w http.ResponseWriter, r *http.Request) {
	token := token.NewToken()
	json.NewEncoder(w).Encode(token.GetBalance())
}

func TransferToken(w http.ResponseWriter, r *http.Request) {
	token := token.NewToken()
	amount := big.NewInt(0)
	token.Transfer(amount)
	json.NewEncoder(w).Encode(token.GetBalance())
}
