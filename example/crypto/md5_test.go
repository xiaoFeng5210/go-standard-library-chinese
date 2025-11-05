package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMD5(t *testing.T) {
	message := "Hello, World!"
	result := generateMD5(message)
	fmt.Println("MD5: ", result)
}

func TestVerifyMD5(t *testing.T) {
	message := "Hello, World!"
	expectedBASE64MD5 := "ZajifYh5KDgxtmS9i38K1A=="
	assert.True(t, verifyMD5(message, expectedBASE64MD5))
}
