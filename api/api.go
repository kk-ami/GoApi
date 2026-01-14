package api

import(
	"encoding/json"
	"net/http"
)

//code balance params
type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	//success code
	Code int
	//account balance
	Balance int64
}

type Error struct {
	//error code
	Code int

	//Error message
	Message string
}