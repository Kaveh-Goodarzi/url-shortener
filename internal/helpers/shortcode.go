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

func CheckDuplicateShortCode(shortCode string) error {
	if _, exists := database.URLstore[shortCode]; exists {
		return errors.New("short code already exists")
	}

	return nil
}
