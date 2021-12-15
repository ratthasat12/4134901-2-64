package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main(){
	fmt.Print("Hello Go ..")
	fmt.Print(time.Now())

	//var num int
	router := mux.NewRouter()

	router.HandleFunc("/",index)

	http.ListenAndServe(":4000",router)
}

func index(res http.ResponseWriter,req *http.Request) {
	fmt.Println("Hello Go..")
	fmt.Println(time.Now())
}