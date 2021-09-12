package main

import (
	"io/ioutil"
	"net/http"
)

type (
	handler struct{}
)

func main() {
	h := &handler{}
	http.ListenAndServe(":8081", h)
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("need POST method"))
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(body)
}
