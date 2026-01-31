package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("%s", fmt.Sprintf("server running on port %s", os.Getenv("PORT")))

	return srv.ListenAndServe()
}
