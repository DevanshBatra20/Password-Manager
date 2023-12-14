package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func EncryptAES(plainText string) (string, error) {
	block, err := aes.NewCipher([]byte("passphrasewhichneedstobe32bytes!"))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	noonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, noonce); err != nil {
		return "", err
	}

	encryptedPassword := gcm.Seal(noonce, noonce, []byte(plainText), nil)

	return base64.StdEncoding.EncodeToString(encryptedPassword), nil
}

func DecryptAES(cipherText string) (string, error) {

	decodedCipherText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte("passphrasewhichneedstobe32bytes!"))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	noonceSize := gcm.NonceSize()
	if len(decodedCipherText) < noonceSize {
		return "", err
	}

	nonce, decodedCipherText := decodedCipherText[:noonceSize], decodedCipherText[noonceSize:]
	plainText, err := gcm.Open(nil, []byte(nonce), decodedCipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), err
}
