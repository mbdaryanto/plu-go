package main

import (
	"fmt"
	"testing"

	"github.com/fernet/fernet-go"
)

func TestFernet(t *testing.T) {
	var originalKey fernet.Key
	if err := (&originalKey).Generate(); err != nil {
		t.Fatalf("Generate error %v", err)
	}
	keyString := originalKey.Encode()
	fmt.Printf("Key: %v\n", keyString)
	key, decodeErr := fernet.DecodeKey(keyString)
	if decodeErr != nil {
		t.Fatalf("DecodeKey error %v", decodeErr)
	}
	// if *originalKey != *key {
	// 	t.Fatal("Key is not equals")
	// }
	plainText := "This is the plain text"
	cipher, encryptErr := fernet.EncryptAndSign([]byte(plainText), key)
	if encryptErr != nil {
		t.Fatalf("Error Encrypting! %v", encryptErr)
	}
	cipherText := string(cipher)
	fmt.Printf("Cipher text: %v\n", string(cipherText))
	msg := fernet.VerifyAndDecrypt([]byte(cipherText), 0, []*fernet.Key{key})
	fmt.Printf("Plain text: %v\n", string(msg))
}
