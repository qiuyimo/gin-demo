package main

import (
	"github.com/qiuyuhome/gin-demo/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouters()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		// ReadHeaderTimeout: 0,
		// TLSConfig:         nil,
		// IdleTimeout:       0,
		// TLSNextProto:      nil,
		// ConnState:         nil,
		// ErrorLog:          nil,
		// BaseContext:       nil,
		// ConnContext:       nil,
	}

	s.ListenAndServe()
}
