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

func SliceTypeToFloat64(numsString []string) ([]float64, error) {
	var numsFloat []float64

	for _, numStr := range numsString {
		numFloat, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return nil, err
		}

		numsFloat = append(numsFloat, numFloat)
	}

	return numsFloat, nil
}
