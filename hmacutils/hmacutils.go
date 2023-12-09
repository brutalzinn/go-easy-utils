package hmacutils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func CalculateSha256HMAC(input []byte, secret []byte) string {
	hmac := hmac.New(sha256.New, secret)
	hmac.Write([]byte(input))
	dataHmac := hmac.Sum(nil)
	hmacHex := hex.EncodeToString(dataHmac)
	return hmacHex
}
