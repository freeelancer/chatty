package cli

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

func GetPemDataBs(filePath string) ([]byte, error) {
	publicKeyFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer publicKeyFile.Close()

	// Read the public key from the PEM file
	pemData, err := io.ReadAll(publicKeyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("Error decoding public key")
	}
	return pemData, nil
}

func GetRSAPubKeyFromFile(filePath string) (*rsa.PublicKey, error) {
	// Read the public key from the PEM file
	pemData, err := GetPemDataBs(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("Error decoding public key")
	}

	rsaPublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Error parsing public key:", err)
	}
	return rsaPublicKey, nil
}

func EncryptMessageWithPubKey(message, filePath string) (string, error) {
	// Load the recipient's public key from a PEM file
	rsaPublicKey, err := GetRSAPubKeyFromFile(filePath)
	if err != nil {
		return "", err
	}

	// Message to be encrypted
	messageBs := []byte(message)
	// Encrypt the message using the recipient's public key
	encryptedMessage, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, messageBs)
	if err != nil {
		return "", err
	}

	// code that decrypts encrypted message
	// privateKeyFile, err := os.Open("/Users/pundix/.ssh/chatty.pem")
	// if err != nil {
	// 	return "", err
	// }
	// defer privateKeyFile.Close()
	// // Read the public key from the PEM file
	// pemData2, err := io.ReadAll(privateKeyFile)
	// if err != nil {
	// 	return "", err
	// }

	// block2, _ := pem.Decode(pemData2)
	// if block2 == nil {
	// 	return "", fmt.Errorf("Error decoding public key")
	// }

	// rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(block2.Bytes)
	// if err != nil {
	// 	return "", fmt.Errorf("Error parsing private key:", err)
	// }

	// messageBs2, err := rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, encryptedMessage)
	// if err != nil {
	// 	return "", err
	// }

	// Print or save the ciphertext (encrypted message)
	return hex.EncodeToString(encryptedMessage), nil
}
