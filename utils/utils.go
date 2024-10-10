package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func GenerateKeys() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	return privateKey, &privateKey.PublicKey, nil
}

func SavePEMKey(fileName string, key *ecdsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	privBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}

	var pemKey = &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: privBytes,
	}

	return pem.Encode(outFile, pemKey)
}

func SavePublicPEMKey(fileName string, pubkey *ecdsa.PublicKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	pubBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return err
	}

	var pemKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}

	return pem.Encode(outFile, pemKey)
}
