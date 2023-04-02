package charlie

import (
	"fmt"
	"io"
	"net/http"

	"github.com/claudiumocanu/go-mtls-example/common"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from charlie\n")
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
	fmt.Fprintf(w, "charlie.pingAlpha(): %s\n", bodyString)
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
	fmt.Fprintf(w, "charlie.pingBravo(): %s\n", bodyString)
}

func StartServer() {
	charlieMux := http.NewServeMux()
	charlieMux.HandleFunc("/hello", hello)
	charlieMux.HandleFunc("/ping-alpha", pingAlpha)
	charlieMux.HandleFunc("/ping-bravo", pingBravo)

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
