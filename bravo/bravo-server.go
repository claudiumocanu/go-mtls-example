package bravo

import (
	"fmt"
	"net/http"
)

const PORT = ":20002"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from bravo\n")
}

func StartServer() {
	bravoMux := http.NewServeMux()
	bravoMux.HandleFunc("/hello", hello)
	fmt.Printf("bravo started: https://localhost%s/hello\n", PORT)
	err := http.ListenAndServeTLS(PORT, "cert/bravo.crt", "cert/bravo.key", bravoMux)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
