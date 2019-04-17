package main

import(
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(whttp.ResponseWriter,r*http.Request){
		log.Println("receivedrequest")
		fmt.Fprintf(w,"HelloDocker!!")
	})
	log.Println("startserver")
	server:=&http.Server{Addr:":8080"}
	if err:=server.ListenAndServe(); err!=nil{
		log.Println(err)
	}
}
