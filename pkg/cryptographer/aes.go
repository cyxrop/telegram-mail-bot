package cryptographer

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type AES struct {
	key []byte
}

func NewAES(key string) *AES {
	return &AES{
		key: []byte(key),
	}
}

func (crypt AES) Encrypt(plaintext string) (string, error) {
	c, err := aes.NewCipher(crypt.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	return encodeBase64(gcm.Seal(nonce, nonce, []byte(plaintext), nil)), nil
}

func (crypt AES) Decrypt(encodedB64 string) (string, error) {
	encoded := decodeBase64(encodedB64)
	c, err := aes.NewCipher(crypt.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(encoded) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := encoded[:nonceSize], encoded[nonceSize:]
	decoded, err := gcm.Open(nil, nonce, ciphertext, nil)

	return string(decoded), err
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}
