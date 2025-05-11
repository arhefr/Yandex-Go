package repeatible

import (
	crypto "crypto/rand"
	"math/big"
	"strconv"
	"time"
)

func NewCryptoRand() string {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(1000000))
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(int(safeNum.Int64()))
}

func DoWithTries(fn func() error, attemps int, delay time.Duration) (err error) {
	for attemps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemps--

			continue
		}

		return nil
	}

	return
}
