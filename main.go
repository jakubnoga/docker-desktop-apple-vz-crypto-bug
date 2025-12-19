package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

var keyHex = "0000000000000000000000000000000000000000000000000000000000000000"

func encryptChallenge(plaintext []byte, aad []byte) (string, error) {
	key, _ := hex.DecodeString(keyHex)
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}

	ciphertext := aead.Seal(nonce, nonce, plaintext, aad)
	return hex.EncodeToString(ciphertext), nil
}

func decryptChallenge(ciphertextHex string, aad []byte) ([]byte, error) {
	key, _ := hex.DecodeString(keyHex)
	msg, _ := hex.DecodeString(ciphertextHex)

	nonce := msg[:chacha20poly1305.NonceSizeX]
	ciphered := msg[chacha20poly1305.NonceSizeX:]

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		panic(err)
	}

	return aead.Open(nil, nonce, ciphered, aad)
}

func testData(plaintext []byte, name string) {
	ciphertextHex, err := encryptChallenge(plaintext, []byte("{}"))
	if err != nil {
		panic(err)
	}

	_, err = decryptChallenge(ciphertextHex, []byte("{}"))
	if err != nil {
		fmt.Printf("%s: FAILED (%d bytes)\n", name, len(plaintext))
	} else {
		fmt.Printf("%s: PASSED (%d bytes)\n", name, len(plaintext))
	}
}

func main() {
	fmt.Println("=== Apple Virtualization Framework Bug ===")
	fmt.Println("XChaCha20Poly1305 decryption fails at >321 bytes")
	fmt.Println()

	// Test under 320 bytes (should pass)
	small := make([]byte, 319)
	rand.Read(small)
	testData(small, "Under threshold (319 bytes)")

	// Test at least 320 bytes (should fail)
	large := make([]byte, 321)
	rand.Read(large)
	testData(large, "At threshold (321 bytes)")
}
