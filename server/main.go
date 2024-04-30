package main

import (
	"fmt"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/{variable}",func(w http.ResponseWriter,r *http.Request){
        variable := r.PathValue("variable")
        fmt.Printf("varaible %v",variable)
    })
    
    mux.HandleFunc("/login/{user}/{pass}",func(w http.ResponseWriter,r *http.Request){
        user := r.PathValue("user")
        pass := r.PathValue("pass")
        fmt.Printf("inicia sesion con %v,%v",user,pass)
    })
    
    if err := http.ListenAndServe("localhost:8080",mux); err != nil{
        fmt.Println(err.Error())
    }

}
