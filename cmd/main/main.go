package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/agarwalson02/go-bookstore/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStores(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9000",r))
	
}