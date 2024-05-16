package main

import (
	"fmt"
    "net/http"
    "database/sql"
  _ "github.com/lib/pq"
    "github.com/rs/cors"   
    "enconding/json"
)

func main() {
    mux := http.NewServeMux()
 
    /*
    ==========================================================================
    =   DATABASE CONNECTION
    ==========================================================================
    */
    db, err := sql.Open("postgres", "postgres://myuser:mypassword@my_postgres/my_database?sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Check connection to db
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}
    fmt.Println("Connected to database")
    
    /*
    ==========================================================================
    =   HANDLERS
    ==========================================================================
    */

    mux.HandleFunc("POST /newUser/", func(w http.ResponseWriter, r *http.Request) {
        type User struct {
            username string `json:"username"`
            password string `json:"password"`
            email string `json:"email"`
        }
        var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

        res, error := db.Exec("INSERT INTO User (name,password,email) VALUES ($1,$2)", user.username,user.password,user.email)
        if error != nil {
			fmt.Printf("Error: %v", error)
		}
		fmt.Printf("Respuesta: %v\n", res)

		// Respond with a success message
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message": "True"}`)
	})

    mux.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
        fmt.Fprintln(w,"Hello World")
    })

    handler := cors.Default().Handler(mux)
	if err := http.ListenAndServe(":9090", handler); err != nil {
		fmt.Println(err.Error())
	}  
}
