package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/Kaveh-Goodarzi/url-shortner/internal/database"
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

func CheckDuplicateShortCode(shortCode string) error {
	if _, exists := database.URLstore[shortCode]; exists {
		return errors.New("short code already exists")
	}

	return nil
}
