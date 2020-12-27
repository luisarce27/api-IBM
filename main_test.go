package main

//Creation of imports
import (

	"testing"
	"net/http"
	"bytes"
	"net/http/httptest"

)

//Creation of unit testing
//Testing if the Encrypt works fine, checking status
func Test_API_Encrypt_Succesfully(t *testing.T){
	body := bytes.NewBufferString(`"sdfsd"`)

	req, err := http.NewRequest("POST", "/encrypt", body)
	
	if err != nil{
		t.Errorf("Error al crear el request")
	}

	recorder := httptest.NewRecorder()

	encrypt(recorder, req)
	
    if recorder.Code != http.StatusOK{
		t.Errorf("Se esperaba status 200. Respuesta %d", recorder.Code)
	}
}
//Tetsing the response in case of a encrypt empty
func Test_API_Encrypt_Empty(t *testing.T){
	body := bytes.NewBufferString(`""`)

	req, err := http.NewRequest("POST", "/encrypt", body)
	
	if err != nil{
		t.Errorf("Error al crear el request")
	}

	recorder := httptest.NewRecorder()

	encrypt(recorder, req)

    if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Se esperaba status 500. Respuesta %d", recorder.Code)
	}
}

//Testing the response in case of dencrypt successfull
func Test_API_Dencrypt_Succesfully(t *testing.T){
	body := bytes.NewBufferString(`"a0455169cad6c6ef3b4c3a2d4c860bb0b7143ca4d86314d873f7a6edde560a3e23f4aa5e2470ec91fa71ca79b4"`)

	req, err := http.NewRequest("POST", "/dencrypt", body)
	
	if err != nil{
		t.Errorf("Error al crear el request")
	}

	recorder := httptest.NewRecorder()

	dencrypt(recorder, req)
	
    if recorder.Code != http.StatusOK{
		t.Errorf("Se esperaba status 200. Respuesta %d", recorder.Code)
	}
}
//Testing the response in case of dencrypt empty
func Test_API_Dencrypt_Empty(t *testing.T){
	body := bytes.NewBufferString(`""`)

	req, err := http.NewRequest("POST", "/dencrypt", body)
	
	if err != nil{
		t.Errorf("Error al crear el request")
	}

	recorder := httptest.NewRecorder()

	dencrypt(recorder, req)

    if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Se esperaba status 500. Respuesta %d", recorder.Code)
	}
}
//Testing the response of dencrypt case of wrong lenght
func Test_API_Dencrypt_WrongLength(t *testing.T){
	body := bytes.NewBufferString(`"a0455169cad6c6ef3b4c3a2d4c860bb0b7143ca4d86314d873f7a6edde560a3e23f4aa5e2470ec91fa71ca79b"`)

	req, err := http.NewRequest("POST", "/dencrypt", body)
	
	if err != nil{
		t.Errorf("Error al crear el request")
	}

	recorder := httptest.NewRecorder()

	dencrypt(recorder, req)

    if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Se esperaba status 500. Respuesta %d", recorder.Code)
	}
}
//Testing the response of dencrypt case of wrong text
func Test_API_Dencrypt_WrongText(t *testing.T){
	body := bytes.NewBufferString(`"a0455169cad6c6ef3b4c3a2d4c860bb0b7143ca4d86314d873f7a6edde560a3e23f4aa5e2470ec91fa71ca79b0"`)

	req, err := http.NewRequest("POST", "/dencrypt", body)
	
	if err != nil{
		t.Errorf("Error al crear el request")
	}

	recorder := httptest.NewRecorder()

	dencrypt(recorder, req)

    if recorder.Code != http.StatusInternalServerError {
		t.Errorf("Se esperaba status 500. Respuesta %d", recorder.Code)
	}
}