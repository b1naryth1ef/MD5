package md5

import (
	"strconv"
	"strings"
)

func parseLots(s string) []float64 {
	l := strings.Split(s, " ")
	res := make([]float64, 0)
	for _, val := range l {
		if (val != " ") && (val != ")") && (val != "(") && (val != "") {
			v, _ := strconv.ParseFloat(val, 64)
			res = append(res, v)
		}
	}
	return res
}
