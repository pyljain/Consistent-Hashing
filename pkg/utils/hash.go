package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
)

const (
	ringSize = 10000
)

func GenerateHash(inputString string) (int, error) {
	hasher := md5.New()
	_, err := io.WriteString(hasher, inputString)
	if err != nil {
		return -1, err
	}

	hashedString := hex.EncodeToString(hasher.Sum(nil))

	// hexToInt, err := strconv.ParseInt(hashedString, 16, 64)
	// if err != nil {
	// 	return -1, err
	// }

	hexToInt, ok := new(big.Int).SetString(hashedString[:16], 16)
	if !ok {
		fmt.Println("Error converting hash to integer")
		return -1, err
	}

	rs := new(big.Int).SetInt64(ringSize)
	mod := new(big.Int).Mod(hexToInt, rs)

	return int(mod.Int64()), nil
}
