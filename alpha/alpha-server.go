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
	fmt.Printf("alpha started: http://localhost%s/hello\n", PORT)
	http.ListenAndServe(PORT, alphaMux)
}
