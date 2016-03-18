package http

import (
	"github.com/open-falcon/recivers/g"
	"github.com/toolkits/file"
	"net/http"
	"strings"
)

func configReciversRoutes() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(g.VERSION))
	})

	http.HandleFunc("/recivers", func(w http.ResponseWriter, r *http.Request) {

	})
}
