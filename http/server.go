package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandler *HTTPHandler
}

func NewHTTPServer(httpHandler *HTTPHandler) *HTTPServer {
	return &HTTPServer{
		httpHandler: httpHandler,
	}
}

func (h *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/employees").Methods("POST").HandlerFunc(h.httpHandler.HandlerCreateWorker)
	router.Path("/employees").Methods("GET").HandlerFunc(h.httpHandler.HandlerGetAllWorkers)
	router.Path("/employees/{id}").Methods("DELETE").HandlerFunc(h.httpHandler.HandlerDeleteWorker)

	fmt.Println("Сервер запущен!")

	err := http.ListenAndServe(":5050", router)

	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
