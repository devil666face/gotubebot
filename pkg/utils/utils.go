package utils

import (
	"strconv"
	// "time"
)

func ToInt64(s interface{}) int64 {
	i, _ := strconv.Atoi(s.(string))
	return int64(i)
}

func ToUint(s interface{}) uint {
	i, _ := strconv.Atoi(s.(string))
	return uint(i)
}

// func GetTimeAgo(t time.Duration) time.Time {
// 	now := time.Now()
// 	// return now.Add(-time.Minute * 15)
// 	return now.Add(t)
// }
