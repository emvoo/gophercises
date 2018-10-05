package normalize

import (
	"strconv"

	)

func Normalize(str string) string {
	var temp string
	for _, v := range str {
		if len(string(v)) != 1 {
			continue
		}
		i, err := strconv.ParseInt(string(v), 10, 32)
		if err != nil {
			continue
		}
		temp += strconv.FormatInt(i, 10)
	}

	return temp
}
