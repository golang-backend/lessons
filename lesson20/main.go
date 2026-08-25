package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	route := http.NewServeMux()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		fmt.Println("Request method: ", method)
		w.Write([]byte("Salom Dunyo"))
	})

	fmt.Println("Server is running!!!")
	if err := http.ListenAndServe(":8180", route); err != nil {
		log.Fatalf("could run sever on port 8180: %v\n", err)
	}

}
