package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/franzego/server-monitor/internal/handlers"
	"github.com/franzego/server-monitor/internal/middleware"
	"github.com/gorilla/mux"
)

func main() {
	//my router
	rou := mux.NewRouter()

	//Registering Routes
	rou.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	rou.HandleFunc("/book", handlers.BookHandler).Methods("GET")

	rou.Use(middleware.LogMiddleware)

	//To Make my server more customizable.
	s := &http.Server{
		Addr:           ":8080",
		Handler:        rou,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//log.Fatal(s.ListenAndServe())

	go func() {
		log.Println("Starting server on :8080")
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("%v error has occured", err)
		}

	}()
	//This is a channel that is on standby to receice an event stream such as a ctrl+c
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// This is to handle graceful shutown. It accounts for both a failure and an Intentional Shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown failed: %v", err)
	}

	log.Println("Server exited cleanly")

}
