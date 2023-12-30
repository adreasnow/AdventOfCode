package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func hashZeroes(s string, n int) string {

	var hexHash string
	var hash [16]byte
	prefix := ""
	for i := 1; i <= n; i++ {
		prefix += "0"
	}

	for i := 0; i <= 999999999; i++ {
		hash = md5.Sum([]byte(s + fmt.Sprintf("%06d", i)))
		hexHash = hex.EncodeToString(hash[:])

		if strings.HasPrefix(hexHash, prefix) {
			return fmt.Sprintf("%d zeros: %06d gives %s", n, i, hex.EncodeToString(hash[:]))
		}
	}
	return ""
}

func main() {
	input := "bgvyzdsv"
	fmt.Println(hashZeroes(input, 5))
	fmt.Println(hashZeroes(input, 6))
}
