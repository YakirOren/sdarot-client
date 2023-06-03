package helpers

import (
	"strconv"
)

func IntToArr(count int) []string {
	arr := make([]string, count)

	for i := 1; i <= count; i++ {
		arr[i-1] = strconv.Itoa(i)
	}
	return arr
}
