package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
)

// SaltPassword increases the ASCII value of each character in the password by 28
func SaltPassword(password string) string {
	saltedPassword := []rune(password)
	for i, char := range saltedPassword {
		saltedPassword[i] = char + 28
	}
	return string(saltedPassword)
}

// HashPassword hashes the salted password using bcrypt
// Encrypt encrypts plain text string into cipher text string using AES algorithm
func Encrypt(plainText string) (string, error) {
	privateKey1 := os.Getenv("PRIVATE_KEY1")
	key := []byte(privateKey1)[:32] // AES-256 requires a 32-byte key

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	privateKey2 := os.Getenv("PRIVATE_KEY2")
	nonce := []byte(privateKey2) // 12 bytes for AES-256-GCM
	cipherText := gcm.Seal(nil, nonce, []byte(plainText), nil)

	return base64.URLEncoding.EncodeToString(cipherText), nil
}
