package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)
type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article
func allArticles(w http.ResponseWriter, r*http.Request){
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},s
	}
	handleRequest()
}