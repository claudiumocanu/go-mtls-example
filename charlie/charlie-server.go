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
	fmt.Println("charlie starting HTTP listener on port", PORT)
	http.ListenAndServe(PORT, charlieMux)
}
