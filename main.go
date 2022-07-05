package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	for i := 1; i < 10; i++ {
				for j := 1; j < i+1; j++ {
					fmt.Print(i, "*", j, "=", i*j)
					fmt.Print("\t")
				}
				fmt.Println()
			}
	io.WriteString(w, "Hello World")

}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", hello)
	log.Print("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}