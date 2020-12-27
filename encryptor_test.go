package main


import (
	"fmt"
    "testing"
)

func Test_Encrypt_Succesfully(t *testing.T){
	var data string = "Texto a encriptar"
	
	response, err := encryptText(data)

	if err != nil {
		t.Errorf("Error en encriptación")
	}

	fmt.Println(response)
}

func Test_Encrypt_Empty(t *testing.T){
	var data string = ""
	
	response, err := encryptText(data)

	if err != nil {
		t.Errorf("Error en encriptación")
	}

	fmt.Println(response)
}

func Test_Dencrypt_Succesfully(t *testing.T){
	var data string = "a0455169cad6c6ef3b4c3a2d4c860bb0b7143ca4d86314d873f7a6edde560a3e23f4aa5e2470ec91fa71ca79b4"
	
	response, err := dencryptText(data)

	if err != nil {
		t.Errorf("Error en desencriptación")
	}

	fmt.Println(response)
}

func Test_Dencrypt_WrongLength(t *testing.T){
	var data string = "a0455169cad6c6ef3b4c3a2d4c860bb0b7143ca4d86314d873f7a6edde560a3e23f4aa5e2470ec91fa71ca79b"
	
	response, err := dencryptText(data)

	if err == nil {
		t.Errorf("Error en desencriptación")
	}

	fmt.Println(response)
}

func Test_Dencrypt_WrongText(t *testing.T){
	var data string = "a0455169cad6c6ef3b4c3a2d4c860bb0b7143ca4d86314d873f7a6edde560a3e23f4aa5e2470ec91fa71ca79b0"
	
	response, err := dencryptText(data)

	if err == nil {
		t.Errorf("Error en desencriptación")
	}

	fmt.Println(response)
}