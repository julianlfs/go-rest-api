package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// People - Estrutura para Pessoa
type People struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}

var Peoples []People

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Peoples)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Peoples {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// obtem o body do POST request
	// para a Struct People
	// adicionando em um array de Peoples
	reqBody, _ := ioutil.ReadAll(r.Body)
	var people People
	json.Unmarshal(reqBody, &people)
	// atualiza a matriz global de Peoples para incluir
	// a nossa nova People
	Peoples = append(Peoples, people)

	json.NewEncoder(w).Encode(people)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]

	for index, article := range Peoples {
		if article.ID == id {
			Peoples = append(Peoples[:index], Peoples[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/peoples", returnAllArticles)
	myRouter.HandleFunc("/people", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/people/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/people/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Peoples = []People{
		People{ID: "1", Name: "Willian Kaminski"},
	}
	handleRequests()
}
