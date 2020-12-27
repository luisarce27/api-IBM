# api-IBM

Golang Encrypt and Dencrypt text
With this project you can encrypt and decrypt a text using an API 

Pre-requisites 📋
- You will need to have installed Golang in your device (https://golang.org/doc/install)
- You need a editor text as VS Code
- You need to have installed Docker


Installation process 🔧

Once having the VS Code installed you have to download extensions of Go and Docker

Unit Testing ⚙️

Once you have downloaded and built the project well you can check via the unit testing the methods of the Api, just running the following code:
(go test --cover)

Example⌨️

Open the Postman and go to create a POST request
Complete the URL with http://localhost:8000/api/encrypt to test encrypt method or http://localhost:8000/api/dencrypt to test the dencrypt method
With the Encrypt
-Put in the body a single text as "test" then press send. You will recive a response like "2f6ce02d93e35030eaed4e5213f65975cd797bcab8e914c4ab67fabb9eaa68f6"

With the Dencrypt 
- Go to Headers and in the Key put Content-Type and in the Value Application/json
- Then in the doby fill with and encrypted string as "2f6ce02d93e35030eaed4e5213f65975cd797bcab8e914c4ab67fabb9eaa68f6" 
- Push send. You will receive and response like "test" 


Built with 🛠️
- Golang 
- Docker
- Postman


Wiki 📖
https://github.com/gorilla/mux


Author ✒️
Luis Arce Picado
