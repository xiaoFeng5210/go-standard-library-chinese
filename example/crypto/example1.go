package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// 返回生成的mac和base64编码的mac
func generateHMAC(secret string, message string) ([]byte, string) {
	hmac := hmac.New(sha256.New, []byte(secret))
	hmac.Write([]byte(message))
	expectedMAC := hmac.Sum(nil)
	fmt.Println("HMAC: ", base64.StdEncoding.EncodeToString(expectedMAC))
	base64Str := base64.StdEncoding.EncodeToString(expectedMAC)
	return expectedMAC, base64Str
}

// 验证
// 第二个参数是明文数据，第三个数据是期待的mac加密数据
func verifyHMAC(secret string, message string, expectedBASE64MAC string) bool {
	messageMAC, _ := generateHMAC(secret, message)
	expectedMAC, _ := base64.StdEncoding.DecodeString(expectedBASE64MAC)
	return hmac.Equal(messageMAC, expectedMAC)

}
