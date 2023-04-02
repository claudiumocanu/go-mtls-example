package alpha

import (
	"fmt"
	"net/http"

	"github.com/claudiumocanu/go-mtls-example/common"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from alpha\n")
}

func StartServer() {
	alphaMux := http.NewServeMux()
	alphaMux.HandleFunc("/hello", hello)
	fmt.Printf("alpha server reachable at: http://%s%s/hello\n",
		common.BaseUrl, common.AlphaServerPort)

	// Clear HTTP expose here
	err := http.ListenAndServe(common.AlphaServerPort, alphaMux)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
