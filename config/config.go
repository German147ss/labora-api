package config

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func StartServer(router http.Handler) error {
	port := ":9000"
	servidor := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Starting alfred on port %s...\n", port)
	if err := servidor.ListenAndServe(); err != nil {
		return fmt.Errorf("Error while starting up alfred: '%v'", err)
	}

	return nil
}
