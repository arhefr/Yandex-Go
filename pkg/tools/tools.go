package tools

import (
	crypto "crypto/rand"
	"math/big"
	"strconv"
)

func NewCryptoRand() string {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(1000000))
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(int(safeNum.Int64()))
}
