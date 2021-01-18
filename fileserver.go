package fileserver

import (
	"fmt"
	"log"
	"net/http"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// Serve start serving static file from given directoy
func Serve(dir string) {
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)
	log.Println("Listening on :18001 ...")
	err := http.ListenAndServe(":18001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
