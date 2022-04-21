package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func wrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := NewLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, r)
		log.Println(r.RemoteAddr, lrw.statusCode)
	})
}

func wrapHandlerWithHeader(wrappedHandler http.Handler) http.Handler {
	version, _ := os.LookupEnv("VERSION")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			w.Header().Add(k, strings.Join(v, ","))
		}
		w.Header().Add("VERSION", version)
		wrappedHandler.ServeHTTP(w, r)
	})
}

func main() {

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: wrapHandlerWithLogging(wrapHandlerWithHeader(http.DefaultServeMux)),
	}

	panic(server.ListenAndServe())
}
