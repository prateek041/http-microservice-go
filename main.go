package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/prateek041/microservices-with-go/handlers"
)

func main() {
	// logger
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	// creating a new ServeMux.
	serveMux := mux.NewRouter() // this is the route router.

	// getting the products handler
	productHandler := handlers.NewProduct(logger)

	getRouter := serveMux.Methods(http.MethodGet).Subrouter() // creating a subrouter that handles only Get requests
	getRouter.HandleFunc("/", productHandler.GetProducts)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter() // Subrouter to handle put requests.
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProduct)
	putRouter.Use(productHandler.MiddlewareTestFunction) // attaching the middleware

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProduct)
	postRouter.Use(productHandler.MiddlewareTestFunction)

	// defining what we want in the server.
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// this is a seperate Go routine
	go func() {
		logger.Println("Starting the server at port 9090")

		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// We have to create a channel that listens to Interrupt and Kill signals
	signalChannel := make(chan os.Signal)      // this create an unbuffered channel because we have not passed any value in here.
	signal.Notify(signalChannel, os.Interrupt) // this listens to Interrupt and Kill signals
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel // this is blocking code. It waits till sigchannel recieves a value.
	logger.Println("Recieved signal, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second) // this creates a context.
	server.Shutdown(tc)

}
