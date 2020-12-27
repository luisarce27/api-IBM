package main

//Creation of imports
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"fmt"
)
//Creation of a security key
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

    
    gcm, err := cipher.NewGCM(c)
    
    // handling errors
    if err != nil {
		return "", err
	}
	
	

    // creation of a new byte array the size of the nonce
    
    nonce := make([]byte, gcm.NonceSize())
    
    // random sequence
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
    }

	//Encryption oftext using the Seal function
	//Using hexadecimal
   	result := gcm.Seal(nonce, nonce, text, nil)
	encryptedText := hex.EncodeToString(result)
	
	fmt.Println(result)
	fmt.Println(encryptedText)

	return encryptedText, nil
}

//Creation of decrytion function
func dencryptText(encryptedText string) (string, error){

	text, err := hex.DecodeString(encryptedText)//Receive the encrypted text 
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

	fmt.Println(rawText)//return the decrypted text (raw text)
	return rawText, nil;
}