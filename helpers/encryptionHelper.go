package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/DevanshBatra20/Password-Manager/configs"
)

var aesSecretKery string = configs.EnvAesSecretKey()

func EncryptAES(plainText string) (string, error) {
	block, err := aes.NewCipher([]byte(aesSecretKery))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	noonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, noonce); err != nil {
		return "", nil
	}

	encryptedPassword := gcm.Seal(noonce, noonce, []byte(plainText), nil)

	return base64.StdEncoding.EncodeToString(encryptedPassword), nil
}

func DecryptAES(cipherText string) (string, error) {
	decodedCipherText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", nil
	}

	block, err := aes.NewCipher([]byte(aesSecretKery))
	if err != nil {
		return "", nil
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", nil
	}

	noonceSize := gcm.NonceSize()
	if len(decodedCipherText) < noonceSize {
		return "", err
	}

	noonce, decodedCipherText := decodedCipherText[:noonceSize], decodedCipherText[noonceSize:]
	plainText, err := gcm.Open(nil, []byte(noonce), decodedCipherText, nil)
	if err != nil {
		return "", nil
	}

	return string(plainText), nil
}
