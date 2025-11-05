package crypto

import (
	"testing"
)

func TestGenerateHMAC(t *testing.T) {
	secret := "1234567890"
	message := "Hello, World!"
	generateHMAC(secret, message)
}

func TestVerifyHMAC(t *testing.T) {
	secret := "1234567890"
	message := "Hello, World!"
	expectedBASE64MAC := "ujePGj7C9YBYv/SB6O4WE8sVwG7Pq+ERjf8ArJLRe1I="
	verifyHMAC(secret, message, expectedBASE64MAC)
}
