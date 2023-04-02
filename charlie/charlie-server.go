package charlie

import (
	"fmt"
	"net/http"

	"github.com/claudiumocanu/go-mtls-example/common"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from charlie\n")
}

func StartServer() {
	charlieMux := http.NewServeMux()
	charlieMux.HandleFunc("/hello", hello)
	fmt.Printf("charlie server reachable at: https://%s%s/hello\n",
		common.BaseUrl, common.CharlieServerPort)

	if err := http.ListenAndServeTLS(
		common.CharlieServerPort,
		"cert/charlie.crt",
		"cert/charlie.key",
		charlieMux,
	); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
