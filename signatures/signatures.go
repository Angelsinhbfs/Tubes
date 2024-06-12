package signatures

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

// SignData signs the given data using the provided private key and returns the signature.
func SignData(data []byte, privateKey *rsa.PrivateKey) (string, error) {
	hashed := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(signature), nil
}

// VerifySignature verifies the given signature using the provided public key and data.
func VerifySignature(data []byte, signature string, publicKey *rsa.PublicKey) error {
	hashed := sha256.Sum256(data)
	sig, err := base64.URLEncoding.DecodeString(signature)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], sig)
}

// SimpleSignature represents a simple signature structure.
type SimpleSignature struct {
	Algorithm string `json:"algorithm"`
	Signature string `json:"signature"`
}

// SignSimpleData signs the given data using the provided private key and returns the simple signature.
func SignSimpleData(data []byte, privateKey *rsa.PrivateKey) (*SimpleSignature, error) {
	signature, err := SignData(data, privateKey)
	if err != nil {
		return nil, err
	}
	return &SimpleSignature{
		Algorithm: "sha256",
		Signature: signature,
	}, nil
}

// VerifySimpleSignature verifies the given simple signature using the provided public key and data.
func VerifySimpleSignature(data []byte, simpleSignature *SimpleSignature, publicKey *rsa.PublicKey) error {
	if simpleSignature.Algorithm != "sha256" {
		return errors.New("unsupported signature algorithm")
	}
	return VerifySignature(data, simpleSignature.Signature, publicKey)
}
