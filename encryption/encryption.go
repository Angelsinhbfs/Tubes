package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

// EncryptedMessage represents the structure of an encrypted message.
type EncryptedMessage struct {
	Encrypted bool   `json:"encrypted"`
	Key       string `json:"key"`
	IV        string `json:"iv"`
	Alg       string `json:"alg"`
	Data      string `json:"data"`
}

// EncryptMessage encrypts the given plaintext message using the recipient's public key.
func EncryptMessage(plaintext []byte, recipientPublicKey *rsa.PublicKey) (*EncryptedMessage, error) {
	// Generate a random AES key and IV
	aesKey := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	// Encrypt the plaintext using AES
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)

	// Encrypt the AES key using the recipient's public key
	encryptedKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, recipientPublicKey, aesKey, nil)
	if err != nil {
		return nil, err
	}

	// Encode the encrypted key, IV, and ciphertext in base64
	encryptedMessage := &EncryptedMessage{
		Encrypted: true,
		Key:       base64.URLEncoding.EncodeToString(encryptedKey),
		IV:        base64.URLEncoding.EncodeToString(iv),
		Alg:       "aes256ctr",
		Data:      base64.URLEncoding.EncodeToString(ciphertext),
	}

	return encryptedMessage, nil
}

// DecryptMessage decrypts the given encrypted message using the recipient's private key.
func DecryptMessage(encryptedMessage *EncryptedMessage, recipientPrivateKey *rsa.PrivateKey) ([]byte, error) {
	// Decode the encrypted key, IV, and ciphertext from base64
	encryptedKey, err := base64.URLEncoding.DecodeString(encryptedMessage.Key)
	if err != nil {
		return nil, err
	}

	iv, err := base64.URLEncoding.DecodeString(encryptedMessage.IV)
	if err != nil {
		return nil, err
	}

	ciphertext, err := base64.URLEncoding.DecodeString(encryptedMessage.Data)
	if err != nil {
		return nil, err
	}

	// Decrypt the AES key using the recipient's private key
	aesKey, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, recipientPrivateKey, encryptedKey, nil)
	if err != nil {
		return nil, err
	}

	// Decrypt the ciphertext using AES
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}
