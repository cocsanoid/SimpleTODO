package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	HTTPHandlers *HTTPHandlers
}

func NewHTTPServer(httphandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		HTTPHandlers: httphandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()
	router.Path("/tasks").Methods("POST").HandlerFunc(s.HTTPHandlers.CreateTaskHandler)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.HTTPHandlers.ListAllTasks)
	router.Path("/tasks/{id}").Methods("GET").HandlerFunc(s.HTTPHandlers.GetTaskById)
	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
