package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"os"

	"github.com/mergermarket/go-pkcs7"
)

func Decrypt(text string) (string, error) {
	key := []byte(os.Getenv("AESKEY_TEXT"))
	cipherText, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		panic("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)

	return string(cipherText), nil
}
