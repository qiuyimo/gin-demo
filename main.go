package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qiuyuhome/gin-demo/global"
	"github.com/qiuyuhome/gin-demo/internal/model"
	"github.com/qiuyuhome/gin-demo/internal/routers"
	"github.com/qiuyuhome/gin-demo/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSettring)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouters()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
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

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSettring)
	if err != nil {
		return err
	}
	return nil
}
