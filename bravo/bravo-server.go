package bravo

import (
	"fmt"
	"net/http"

	"github.com/claudiumocanu/go-mtls-example/common"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from bravo\n")
}

func StartServer() {
	bravoMux := http.NewServeMux()
	bravoMux.HandleFunc("/hello", hello)
	fmt.Printf("bravo server reachable at: https://%s%s/hello\n",
		common.BaseUrl, common.BravoServerPort)

	if err := http.ListenAndServeTLS(
		common.BravoServerPort,
		"cert/bravo.crt",
		"cert/bravo.key",
		bravoMux,
	); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
