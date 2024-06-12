package identity

import (
	"Tubes/signatures"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
)

// Identity represents a Nomad identity.
type Identity struct {
	ID        string
	PublicKey *rsa.PublicKey
}

// VerifyIdentity verifies the identity using the provided signature and public key.
func VerifyIdentity(identity *Identity, signature string) error {
	data := []byte(identity.ID)
	return signatures.VerifySignature(data, signature, identity.PublicKey)
}

// LinkIdentities links two identities by verifying their signatures.
func LinkIdentities(identity1, identity2 *Identity, signature1, signature2 string) error {
	err := VerifyIdentity(identity1, signature1)
	if err != nil {
		return err
	}
	return VerifyIdentity(identity2, signature2)
}

// GeneratePortableID generates a portable ID for the given identity.
func GeneratePortableID(identity *Identity) (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(identity.PublicKey)
	if err != nil {
		return "", err
	}
	data := append([]byte(identity.ID), publicKeyBytes...)
	hashed := sha256.Sum256(data)
	return base64.URLEncoding.EncodeToString(hashed[:]), nil
}
