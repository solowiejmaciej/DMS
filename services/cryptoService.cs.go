package services

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
)

func Decrypt(encrypted string) (string, error) {
	var key = []byte(os.Getenv("AES_SECRET"))
	var iv = []byte(os.Getenv("AES_IV"))
	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	mode.CryptBlocks(decrypted, ciphertext)

	decrypted, err = removePKCS7Padding(decrypted)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

func removePKCS7Padding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("removePKCS7Padding: input data is empty")
	}

	pad := data[length-1]
	padLength := int(pad)

	if padLength > length || padLength == 0 {
		return nil, fmt.Errorf("removePKCS7Padding: invalid padding length")
	}

	for _, v := range data[length-padLength:] {
		if v != pad {
			return nil, fmt.Errorf("removePKCS7Padding: invalid padding content")
		}
	}

	return data[:length-padLength], nil
}
