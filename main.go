//Filename: main.go

package main

import (
	"log"
	"net/http"
)

//let's create a function to define middleware function that loops the reuest and response

func LogMiddleware(CartoonNetwork http.Handler) http.Handler { //takes an incoming 'http.Handler argument and returns a new 'http.Handler'
	return http.HandlerFunc(func(icebear http.ResponseWriter, panda *http.Request) {
		//Executed on the way down to the handler
		log.Printf("Executing Middleware Bear")
		//pass the request to the next handler in the chain
		//handles the request
		CartoonNetwork.ServeHTTP(icebear, panda)
		//Executed on the way up to the client
		log.Printf("Executing Middleware Bear again")

	})
}

func Middleware2(CartoonNetwork http.Handler) http.Handler {
	return http.HandlerFunc(func(icebear http.ResponseWriter, panda *http.Request) {
		//Executed on the way down to the handler
		log.Println("We bare Bears 2")
		if panda.URL.Path == "/channel" {
			return
		}
		CartoonNetwork.ServeHTTP(icebear, panda)
		//Executed on the way up to the client
		log.Println("We bare Bears 2 again")

	})
}

func finalhandler(mordecai http.ResponseWriter, rigby *http.Request) {

	//Log a message to indicate that the handler is being executed
	log.Println("Regular Show ")
	//Write the string "You should watch Regualr Show"
	mordecai.Write([]byte("You should watch 'Regular Show' "))
}

func main() {
	//Create a new multiplexer (router)
	mux := http.NewServeMux()
	//Register a handler function for the root path "/" that will go through the middleware chain
	mux.Handle("/", LogMiddleware(Middleware2(http.HandlerFunc(finalhandler))))

	//log a message indicating that the server is starting
	log.Print("Server :8282")
	//Start the server on port 8282 and use the multiplexer to handle incoming requests
	err := http.ListenAndServe(":8282", mux)
	//if there was an error starting the server, log the error and exit the program
	log.Fatal(err)
}
