package server

import (
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/viniciusidacruz/hexagonal-archtecture/adapters/web/handler"
	"github.com/viniciusidacruz/hexagonal-archtecture/application"
)

type WebServer struct {
	ProductService application.ProductServiceInterface
}

func NewWebServer(ps application.ProductServiceInterface) *WebServer {
	return &WebServer{ProductService: ps}
}

func (w *WebServer) Serve() {
	router := mux.NewRouter()
	handler.MakeProductHandlers(router, negroni.New(), w.ProductService)

	n := negroni.Classic() // logger + recovery + static
	n.UseHandler(router)   // conecta o mux ao negroni

	server := &http.Server{
		Addr:              ":8080",
		Handler:           n,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
