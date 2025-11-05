package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func generateHMAC(secret string, message string) []byte {
	hmac := hmac.New(sha256.New, []byte(secret))
	hmac.Write([]byte(message))
	expectedMAC := hmac.Sum(nil)
	fmt.Println("HMAC: ", base64.StdEncoding.EncodeToString(expectedMAC))
	return expectedMAC
}
