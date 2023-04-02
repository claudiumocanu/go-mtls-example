package alpha

import (
	"fmt"
	"io"
	"net/http"

	"github.com/claudiumocanu/go-mtls-example/common"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from alpha\n")
}

func pingBravo(w http.ResponseWriter, req *http.Request) {
	c := http.Client{}
	res, err := c.Get(fmt.Sprintf("https://%s%s/hello", common.BaseUrl, common.BravoServerPort))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bodyString := string(bodyBytes)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "alpha.pingBravo(): %s\n", bodyString)
}

func pingCharlie(w http.ResponseWriter, req *http.Request) {
	c := http.Client{}
	res, err := c.Get(fmt.Sprintf("https://%s%s/hello", common.BaseUrl, common.CharlieServerPort))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bodyString := string(bodyBytes)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "alpha.pingCharlie(): %s\n", bodyString)
}

func StartServer() {
	alphaMux := http.NewServeMux()

	alphaMux.HandleFunc("/hello", hello)
	alphaMux.HandleFunc("/ping-bravo", pingBravo)
	alphaMux.HandleFunc("/ping-charlie", pingCharlie)

	fmt.Printf("alpha server reachable at: http://%s%s/hello\n",
		common.BaseUrl, common.AlphaServerPort)

	// Clear HTTP expose here
	err := http.ListenAndServe(common.AlphaServerPort, alphaMux)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
