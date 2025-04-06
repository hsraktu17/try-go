package main

import (
	"fmt"
	"io"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, "Get method called")
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintf(w, "This is the About Page!")
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusCreated) // 201
		fmt.Fprintf(w, "You posted to About!")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		fmt.Fprintf(w, "Method not allowed")
	}
}

func handleWorking(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Error reading request body", err)
			return
		}

		defer r.Body.Close()

	}
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/working", handleWorking)
	fmt.Println("Hello starting backend apis")
	fmt.Println("Server working on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
