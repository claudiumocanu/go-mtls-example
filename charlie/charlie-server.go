package charlie

import (
	"fmt"
	"net/http"
)

const PORT = ":20003"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from charlie\n")
}

func StartServer() {
	charlieMux := http.NewServeMux()
	charlieMux.HandleFunc("/hello", hello)
	fmt.Printf("charlie started: https://localhost%s/hello\n", PORT)
	err := http.ListenAndServeTLS(PORT, "cert/charlie.crt", "cert/charlie.key", charlieMux)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
