package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	message := fmt.Sprintf("Hello I'm %s. I'm %s.", name, age)

	w.Write([]byte(message))
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/app/myfamily/family.txt")
	if err != nil {
		log.Fatal("Error reading file", err)
	}
	message := fmt.Sprintf("My family: %s.", string(data))

	w.Write([]byte(message))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	message := fmt.Sprintf("User: %s. Password: %s", user, password)

	w.Write([]byte(message))
}
