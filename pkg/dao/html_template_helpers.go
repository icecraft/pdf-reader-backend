package dao

import (
	"strconv"
	"strings"
)

// make sure that len(ids) > 0 !!
func RenderArrInt(arr []int) string {
	str := make([]string, len(arr))
	for i, v := range arr {
		str[i] = strconv.Itoa(v)
	}
	return strings.Join(str, ", ")
}
