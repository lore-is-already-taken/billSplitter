package main

import (
	"fmt"
    "net/http"
    "database/sql"
  _ "github.com/lib/pq"
)

const(
    host = "mypostgres"
    port = 5432
    user = "myuser"
    password = "mypassword"
    dbname = "mydatabase"
)

func connect(){
    // CONNECTS TO DB
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
    panic(err)
    }
    defer db.Close()
}

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
    
    mux.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
        fmt.Fprintln(w,"Hello World")
    })
    
    mux.HandleFunc("/connect",func(w http.ResponseWriter,r *http.Request){
        connect()
    })
    
    if err := http.ListenAndServe(":9090",mux); err != nil{
        fmt.Println(err.Error())
    }

}
