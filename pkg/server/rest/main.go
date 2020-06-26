package rest

import (
	"context"
	"demo/pkg/services/products"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

type HTTPRest interface {
	Serve()
}

type httpRest struct {
	addr           string
	productService products.Service
}

func NewHTTPRest(addr string, productService products.Service) HTTPRest {
	return &httpRest{
		addr:           addr,
		productService: productService,
	}
}

func (h *httpRest) Serve() {
	chi := chi.NewMux()

	srv := &http.Server{
		Addr:         h.addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      chi,
	}

	// endpoints
	productsEndpoint := newProductsEndpoint(h.productService)
	chi.Mount("/api/products", productsEndpoint.routers())

	// start
	log.Println("Listen on", h.addr)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(errors.Wrap(err, "listen and serve"))
		}
	}()

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(errors.Wrap(err, "graceful shutdown"))
	}
	log.Println("shutting down")
	os.Exit(0)
}
