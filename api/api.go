package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
    Message string `json:"message"`
}

// Если функция начинается с маленькой буквы, она не экспортируется и доступна только внутри пакета
// Если с большой - то экспортируется и доступна в других пакетах
func Api() {
	fmt.Println("API started")

	http.HandleFunc("GET /ebal", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Ebal API is working!")
		fmt.Printf("WOWAN - %v,\n\n RORA - %v",w, r)
		// return {"kek": "lol"}

		response := Response{Message: "Hello, World!"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
	})

	 http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        response := Response{Message: "Hello, World!"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    })

	println("Server started on http://localhost:8080")
    println("GET Endpoints: /hello; /ping")
    http.ListenAndServe(":8080", nil)

}