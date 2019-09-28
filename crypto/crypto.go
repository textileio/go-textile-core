package crypto

import (
	"fmt"

	"github.com/textileio/go-textile-core/crypto/asymmetric"
	"github.com/textileio/go-textile-core/crypto/symmetric"
)

type EncryptionKey interface {
	Encrypt([]byte) ([]byte, error)
	Marshal() ([]byte, error)
}

type DecryptionKey interface {
	EncryptionKey
	Decrypt([]byte) ([]byte, error)
}

// ParseEncryptionKey
func ParseEncryptionKey(k []byte) (EncryptionKey, error) {
	aek, err := asymmetric.NewEncryptionKey(k)
	if err == nil {
		return aek, nil
	}
	sk, err := symmetric.NewKey(k)
	if err == nil {
		return sk, nil
	}

	return nil, fmt.Errorf("parse encryption key failed")
}

func ParseDecryptionKey(k []byte) (DecryptionKey, error) {
	adk, err := asymmetric.NewDecryptionKey(k)
	if err == nil {
		return adk, nil
	}
	sk, err := symmetric.NewKey(k)
	if err == nil {
		return sk, nil
	}

	return nil, fmt.Errorf("parse decryption key failed")
}
