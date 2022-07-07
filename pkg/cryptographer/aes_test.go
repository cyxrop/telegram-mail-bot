package cryptographer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncryptDecrypt(t *testing.T) {
	testCases := []string{
		"",
		"valid",
		"looooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong",
		"4un24gunu359!#@$%^&*(*)()$?><\"{|}",
	}

	crypt := NewAES("the-key-has-to-be-32-bytes-long!")
	for _, str := range testCases {
		t.Run(fmt.Sprintf("encrypt, decrypt %q", str), func(t *testing.T) {
			encrypted, err := crypt.Encrypt(str)
			require.NoError(t, err)

			decrypted, err := crypt.Decrypt(encrypted)
			require.NoError(t, err)

			assert.Equal(t, str, decrypted)
		})
	}
}
