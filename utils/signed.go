package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

// Signed the interface of signed
type Signed struct {
	secretKey string
}

// NewSigned ...
func NewSigned(secretKey string) *Signed {
	s := new(Signed)
	s.secretKey = fmt.Sprintf("%s", sha256.Sum256([]byte(secretKey)))
	return s

}

// AESEncode encrypt the string
func (s *Signed) AESEncode(t string) string {
	plaintext := []byte(t)

	block, err := aes.NewCipher([]byte(s.secretKey))
	if err != nil {
		panic(err)

	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)

	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return fmt.Sprintf("%x", ciphertext)

}

// AESDecode decode the string
func (s *Signed) AESDecode(t string) string {
	ciphertext, _ := hex.DecodeString(t)

	block, err := aes.NewCipher([]byte(s.secretKey))
	if err != nil {
		panic(err)

	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")

	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)

}
