package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"

	dictionary "dict/dictionary"
)

var dict = dictionary.New()	

type RequestAdd struct {
	Word       string
	Definition string
}

type RequestDelete struct {
	Word string
}

type ResponseList struct {
	entries map[string]dictionary.Entry
}

func add(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/add" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch req.Method {
	case "POST":
		var request RequestAdd
		req.ParseForm()
		json.NewDecoder(req.Body).Decode(&request)
		dict.Add(request.Word, request.Definition)

	default:
		fmt.Fprintf(w, "Sorry, only POST method is supported.")
	}
	defer req.Body.Close()
}

func remove(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/remove" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch req.Method {
	case "POST":
		var request RequestDelete
		req.ParseForm()
		json.NewDecoder(req.Body).Decode(&request)
		dict.Remove(request.Word)
	default:
		fmt.Fprintf(w, "Sorry, only GET method is supported.")
	}
	defer req.Body.Close()
}

func list(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/list" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch req.Method {
	case "GET":
		_, entries := dict.List()

		w.Header().Set("Content-Type", "application/json")
		response := make(map[string]dictionary.Entry)
		for key, entry := range entries {
			fmt.Println(entry)
			response[key] = entry
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonResponse)

	default:
		fmt.Fprintf(w, "Sorry, only GET method is supported.")
	}

}

func main() {

	http.HandleFunc("/add", add)
	http.HandleFunc("/remove", remove)
	http.HandleFunc("/list", list)

	http.ListenAndServe(":8090", nil)
}
