package binary

import (
	"strconv"
)

func IsBinary(s string) bool {
	_, err := strconv.ParseUint(s, 2, 32)
	if err != nil {
		return false
	}

	return true
}

func ToBinary(s string) (string, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(int64(num), 2), nil
}
