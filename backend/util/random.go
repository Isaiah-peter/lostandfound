package util

import (
	"math/rand"
	"strings"
	"time"
)

func RandomWord(num int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12345678890"

	var result []string

	rand.Seed(time.Now().UnixNano())
    min := 0
    max := 62

	for i := 0; i < num; i++ {
		num := rand.Intn(max - min) + min
		result = append(result, string(str[num]))
	}
	return strings.Join(result, "")
}