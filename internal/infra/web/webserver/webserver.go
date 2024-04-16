package webserver

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type WebServer struct {
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: webServerPort,
	}
}

func (w *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	w.Handlers[path] = handler
}

func (w *WebServer) Start() {
	mux := http.NewServeMux()

	for path, handler := range w.Handlers {
		mux.Handle(path, LoggingMiddleware(handler))
	}

	err := http.ListenAndServe(":"+w.WebServerPort, mux)
	if err != nil {
		panic(err)
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Println(fmt.Sprintf("%s %s %v", r.Method, r.URL.Path, time.Since(start)))
	})
}
