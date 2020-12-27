package main

//Creation of imports
import (
	"encoding/json"
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

//Post of encrypt text
func encrypt(w http.ResponseWriter, r *http.Request) {
	var data string

	json.NewDecoder(r.Body).Decode(&data)

	if data == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Texto invalido, verifique el ingreso del texto")
		return
	}

	response, err := encryptText(data)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Ha ocurrido un error mientras se encriptaba el texto")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}



//Post of dencrypt text
func dencrypt(w http.ResponseWriter, r *http.Request) {
	var data string
	
	json.NewDecoder(r.Body).Decode(&data)

	if data == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Texto invalido, verifique el ingreso del texto")
		return
	}

	response, err := dencryptText(data)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Ha ocurrido un error mientras se desencriptaba el texto")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func main() {
	fmt.Println("IBM Example API")

	//Init Router
	r := mux.NewRouter()

	// Route Handles/ Endpoints
	r.HandleFunc("/api/encrypt", encrypt).Methods("POST")
	r.HandleFunc("/api/dencrypt", dencrypt).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
