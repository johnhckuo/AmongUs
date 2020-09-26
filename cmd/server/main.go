package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {

	// starting server

		mux := http.NewServeMux()

		mux.HandleFunc("/", homeHandler)
		//mux.HandleFunc("/contact", contactHandler)

		log.Println(fmt.Sprintf("Server running on http://localhost%s 🐹", ":4000"))
		err := http.ListenAndServe(":4000", mux)
		if err != nil {
			log.Fatalf("could not run the server %v", err)
			return
		}

	//handler()
}

// listen to request

func homeHandler(w http.ResponseWriter, r *http.Request) {

	// validation
	w.Write([]byte("Hi"))

}


func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from contact handler"))
}
