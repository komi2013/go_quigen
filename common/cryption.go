package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	// "fmt"
    "io"
	"encoding/hex"
	"encoding/base64"
)

func Encrypt(keyStr string, text string) string {
	key, _ := hex.DecodeString(keyStr)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
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
	enctxt := base64.StdEncoding.EncodeToString(ciphertext) 
	return enctxt
}

func Decrypt(keyStr string, text string) string{
	key, _ := hex.DecodeString(keyStr)
	ciphertext, _ := base64.StdEncoding.DecodeString(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		// panic("ciphertext too short")
		return ""
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	s := string(ciphertext)
	return s
}