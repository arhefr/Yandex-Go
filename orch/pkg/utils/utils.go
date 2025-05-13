package repeatible

import (
	crypto "crypto/rand"
	"math/big"
	"strconv"
	"time"
)

func NewCryptoRand(size int64) string {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(size))
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
