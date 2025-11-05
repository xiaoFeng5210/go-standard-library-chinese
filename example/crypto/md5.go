package crypto

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
)

func generateMD5(message string) string {
	md5 := md5.New()
	md5.Write([]byte(message))
	expectedMD5 := md5.Sum(nil)
	return base64.StdEncoding.EncodeToString(expectedMD5)
}

func verifyMD5(message string, expectedBASE64MD5 string) bool {
	expectedMD5, _ := base64.StdEncoding.DecodeString(expectedBASE64MD5)
	md5 := md5.New()
	md5.Write([]byte(message))
	md5Result := md5.Sum(nil)
	return bytes.Equal(md5Result, expectedMD5)
}
