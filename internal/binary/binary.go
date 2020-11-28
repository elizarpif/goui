package binary

import (
	"strconv"
)

// GF(256)
func IsBinary(s string) bool {
	res, err := strconv.ParseUint(s, 2, 32)
	if err != nil || res > 256 {
		return false
	}

	return true
}

func ToBinary(s string) (int, error) {
	res, err := strconv.ParseInt(s, 2, 32)
	if err != nil {
		return 0, err
	}

	return int(res) % 256, err
}

func ToBinaryStr(n int) string {
	return strconv.FormatInt(int64(n), 2)
}
