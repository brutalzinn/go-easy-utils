package hmacutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// / IN AGREE WITH https://www.freeformatter.com/hmac-generator.html#before-output
func TestValidHmac(t *testing.T) {
	result := "ea286b8ef85441af1e8d9d06fcda2488fa94c1c66319019663f179322eaf9175"
	input := []byte("teste")
	secret := []byte("cf16ce65d721a905bd03eed237591695aae5cc3cc4e632a54488c77dcecdd1fb")
	hmac := CalculateSha256HMAC(input, secret)
	assert.True(t, hmac == result, "hmac invalid")
}
