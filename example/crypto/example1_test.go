package crypto

import (
	"testing"
)

func TestGenerateHMAC(t *testing.T) {
	secret := "1234567890"
	message := "Hello, World!"
	generateHMAC(secret, message)
}
