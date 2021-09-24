package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yogeshkaushik1904/search-engine/controller"
	router "github.com/yogeshkaushik1904/search-engine/http"
	"github.com/yogeshkaushik1904/search-engine/repository"
	"github.com/yogeshkaushik1904/search-engine/routes"
	"github.com/yogeshkaushik1904/search-engine/service"
)

var (
	elasticRepo repository.OrderRepository = repository.NewElasticOrderRepository()
	svc         service.OrderService       = service.NewOrderService(elasticRepo)
	httpRouter  router.Router              = router.NewMuxRouter()
	cntrl       controller.OrderController = controller.NewOrderController(svc)
)

func main() {

	l := log.New(os.Stdout, "search-engine", log.LstdFlags)
	routes.RegisterRoutes(httpRouter, cntrl)
	s := http.Server{
		Addr:         "127.0.0.1:8090",  // configure the bind address
		Handler:      httpRouter,        // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 8090")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
