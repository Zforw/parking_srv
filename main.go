package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"parking/handler"
	"parking/initialize"
	"syscall"
	"time"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
	initialize.InitValidator()
	Router := initialize.Routers()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 35491),
		Handler: Router,
	}
	zap.S().Debugf("启动端口：%d", 35491)
	go func() {
		for {
			time.Sleep(time.Hour)
			handler.TOKEN = handler.GetAccessToken()
		}
	}()
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	re := <-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Fatal(err)
	}
	zap.S().Info("已停止， 监听信号: ", re)
}
