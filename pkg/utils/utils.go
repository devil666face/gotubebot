package utils

import (
	"strconv"
)

func ToInt64(s interface{}) int64 {
	i, _ := strconv.Atoi(s.(string))
	return int64(i)
}

func ToUint(s interface{}) uint {
	i, _ := strconv.Atoi(s.(string))
	return uint(i)
}
