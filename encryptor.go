package main


import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"fmt"
)

var key []byte = []byte("ClaveseguridadparaejemplodeIBM..")

func encryptText(rawText string) (string, error){
    text := []byte(rawText)
	
	fmt.Println(rawText)
	fmt.Println(text)

    //Key creation
    c, err := aes.NewCipher(key)
    //Handling errors
    if err != nil {
		return "", err
    }

    // gcm or Galois/Counter Mode, is a mode of operation
    gcm, err := cipher.NewGCM(c)
    
    // handling errors
    if err != nil {
		return "", err
	}
	
	//out := make([]byte, len(rawText))

    // creates a new byte array the size of the nonce
    // which must be passed to Seal
    nonce := make([]byte, gcm.NonceSize())
    // populates our nonce with a cryptographically secure
    // random sequence
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
    }

    // here we encrypt our text using the Seal function
    // Seal encrypts and authenticates plaintext, authenticates the
    // additional data and appends the result to dst, returning the updated
    // slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	result := gcm.Seal(nonce, nonce, text, nil)
	encryptedText := hex.EncodeToString(result)
	
	fmt.Println(result)
	fmt.Println(encryptedText)

	return encryptedText, nil
}

func dencryptText(encryptedText string) (string, error){

	text, err := hex.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	fmt.Println(encryptedText)
	fmt.Println(text)
	
    c, err := aes.NewCipher(key)
    if err != nil {
		return "", err
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
		return "", err
    }

    nonceSize := gcm.NonceSize()
    if len(text) < nonceSize {
		return "", err
    }

    nonce, text := text[:nonceSize], text[nonceSize:]
    decryptedText, err := gcm.Open(nil, nonce, text, nil)
    if err != nil {
		return "", err
	}
	
	rawText := string(decryptedText)

	fmt.Println(rawText)
	return rawText, nil;
}