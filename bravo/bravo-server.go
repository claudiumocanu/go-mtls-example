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
	fmt.Println("bravo starting HTTP listener on port", PORT)
	http.ListenAndServe(PORT, bravoMux)
}
