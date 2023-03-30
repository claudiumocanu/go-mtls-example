package alpha

import (
	"fmt"
	"net/http"
)

const PORT = ":20001"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from alpha\n")
}

func StartServer() {
	alphaMux := http.NewServeMux()
	alphaMux.HandleFunc("/hello", hello)
	fmt.Println("alpha starting HTTP listener on port", PORT)
	http.ListenAndServe(PORT, alphaMux)
}
