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

    if err := http.ListenAndServe("localhost:8080",mux); err != nil{
        fmt.Println(err.Error())
    }

}
