package engine

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

//var KEY = []byte(GenerateRandomString()) // 32 octets
// EncryptText chiffre un texte en utilisant l'algorithme AES-256

func EncryptText(plaintext []byte, key []byte) (string, error) {
	// Générer une clé de 32 octets pour l'algorithme AES-256

	//Generate a random string of 32 bytes

	//key := KEY
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Générer un vecteur d'initialisation de 16 octets
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Chiffrer le texte brut en utilisant la clé et le vecteur d'initialisation
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)

	// Concaténer le vecteur d'initialisation et le texte chiffré
	result := hex.EncodeToString(iv) + hex.EncodeToString(ciphertext)
	return result, nil
}

// Generate a random string of 32 bytes
func GenerateRandomString() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// DecryptText déchiffre un texte chiffré en utilisant l'algorithme AES-256
func DecryptText(ciphertext string, key []byte) (string, error) {
	// Extraire le vecteur d'initialisation et le texte chiffré
	iv, err := hex.DecodeString(ciphertext[:32])
	if err != nil {
		return "", err
	}
	Text, err := hex.DecodeString(ciphertext[32:])
	if err != nil {
		return "", err
	}

	// Générer une clé de 32 octets pour l'algorithme AES-256
	// var KEY = []byte(GenerateRandomString()) // 32 octets

	//key := KEY
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Déchiffrer le texte chiffré en utilisant la clé et le vecteur d'initialisation
	plaintext := make([]byte, len(string(Text)))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, Text)

	return string(plaintext), nil
}
