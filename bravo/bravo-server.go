package bravo

import (
	"fmt"
	"io"
	"net/http"

	"github.com/claudiumocanu/go-mtls-example/common"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from bravo\n")
}

func pingAlpha(w http.ResponseWriter, req *http.Request) {
	c := http.Client{}
	res, err := c.Get(fmt.Sprintf("http://%s%s/hello", common.BaseUrl, common.AlphaServerPort))
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
	fmt.Fprintf(w, "bravo.pingAlpha(): %s\n", bodyString)
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
	fmt.Fprintf(w, "bravo.pingCharlie(): %s\n", bodyString)
}

func StartServer() {
	bravoMux := http.NewServeMux()
	bravoMux.HandleFunc("/hello", hello)
	bravoMux.HandleFunc("/ping-alpha", pingAlpha)
	bravoMux.HandleFunc("/ping-charlie", pingCharlie)

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
