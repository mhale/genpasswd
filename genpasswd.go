package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var defaultLength = 12 // Default number of chars
var chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var lenChars = len(chars)

func generate(length int) string {
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = chars[r.Intn(lenChars)]
	}
	return string(password)
}

func main() {
	length := defaultLength
	if len(os.Args) > 1 {
		length, _ = strconv.Atoi(os.Args[1])
		if length < 1 {
			length = defaultLength
		}
	}

	password := generate(length)
	fmt.Println(password)
}
