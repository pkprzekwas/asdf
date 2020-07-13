package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkprzekwas/asdf/linked"
)

func main() {
	todos := linked.CreateLL()
	todos.Add("prepare")
	todos.Add("test")
	todos.Add("exec")
	fmt.Println(todos)

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Todos %s\n", todos)
		case "POST":
			data, err := ioutil.ReadAll(r.Body)

			if err != nil {
				fmt.Fprintln(w, "Something went wrong")
				return
			}

			todos.Add(string(data))
			fmt.Fprintf(w, "Todos %s\n", todos)
		default:
			fmt.Fprintln(w, "Method not supported")
		}
	})

	log.Print("Starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
