package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const limit = 15

/*
New ...
	addr := ":8080"

	r := mux.NewRouter()
	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		// web socket connect
		_, e := conn.Accept(res, req)
		if e != nil {
			log.File("accept")
			return
		}

		// or http server

	})

	server.New(addr, r)
*/
func New(
	addr string,
	router *mux.Router,
) {

	s := &http.Server{
		Addr:         addr,
		Handler:      router,
		WriteTimeout: limit * time.Second,
		ReadTimeout:  limit * time.Second,
	}

	go func() {

		e := s.ListenAndServe()
		if e != nil && e != http.ErrServerClosed {
			panic(e)
		}

	}()
}
