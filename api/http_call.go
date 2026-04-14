package api

import (
	"fmt"
	"io"
	"net/http"
)

func MyFirstHttpCall() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Ошибка при чтении ответа: %v\n", err)
        return
    }
	
	fmt.Printf("RES: %v \n\n ERR: %v", string(body), err)
}