package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)
type Article struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

var articles Articles = Articles{
	Article{Id: "2", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	Article{Id: "1", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}
func allArticles(w http.ResponseWriter, r*http.Request){
	fmt.Println("Endpoint Hit: All Articles Endpoint")

	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	fmt.Println("Getting article ")
    vars := mux.Vars(r)
	key := vars["id"]
	for _, article := range articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createNewArticle(w http.ResponseWriter, r *http.Request) { 
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    articles = append(articles, article)

    json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	
    vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Deletin article ")
    for index, article := range articles {
        if article.Id == id {
            articles = append(articles[:index], articles[index+1:]...)
        }
    }

}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", allArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	
	handleRequest()
}