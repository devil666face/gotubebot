package utils

import (
	"strconv"
)

func ToInt64(s interface{}) int64 {
	i, _ := strconv.Atoi(s.(string))
	return int64(i)
}
