package helpers

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateShortCode() string {
	b := make([]byte, 6)
	rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)[:6]
}

var ID = 1
var proccess = 0

func IDGenerator() int {
	if proccess == 0 {
		proccess++
		return ID
	}

	ID++
	proccess++
	return ID
}
