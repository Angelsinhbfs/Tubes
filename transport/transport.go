package transport

import (
	"Tubes/signatures"
	"bytes"
	"crypto/rsa"
	//"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// SignedRequest represents a signed HTTP request.
type SignedRequest struct {
	Method    string
	URL       string
	Headers   map[string]string
	Body      []byte
	Signature string
}

// SendSignedRequest sends a signed HTTP request.
func SendSignedRequest(req *SignedRequest, privateKey *rsa.PrivateKey) (*http.Response, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	// Create the HTTP request
	httpReq, err := http.NewRequest(req.Method, req.URL, bytes.NewBuffer(req.Body))
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	// Sign the request
	signature, err := signatures.SignData(req.Body, privateKey)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Signature", signature)

	// Send the request
	return client.Do(httpReq)
}

// VerifySignedRequest verifies the signature of an HTTP request.
func VerifySignedRequest(req *http.Request, publicKey *rsa.PublicKey) error {
	signature := req.Header.Get("Signature")
	if signature == "" {
		return errors.New("missing signature")
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	return signatures.VerifySignature(body, signature, publicKey)
}
