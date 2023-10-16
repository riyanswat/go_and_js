package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func randomPassword(length int) {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := strings.ToUpper(lower)
	nums := "1234567890"
	syms := "!@#$%^&*"
	allChars := lower + upper + nums + syms

	password := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range password {
		password[i] = allChars[rand.Intn(len(allChars))]
	}

	fmt.Println(string(password))
}

func main() {
	randomPassword(8)
}
