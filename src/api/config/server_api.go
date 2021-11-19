package config

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"proyectos/src/api/config/settings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

//SRV type
type SRV struct {
	*gin.Engine
}

//NewServer Returns gin engine instance
func NewServer(g *gin.Engine) *SRV {
	//Create server
	return &SRV{Engine: g}
}

//Run server
func (r *SRV) Run(cfg *settings.Data, options ...interface{}) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: r,
	}

	//Run server in goroutine
	go func() {
		// service connections
		log.Printf("[package:config][function:SRV.Run] Running server at port : %s ", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with 5 seconds timeout
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// Quit
	log.Println("[package:config][function:SRV.Run] Quit server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println(fmt.Sprintf("[package:config][function:SRV.Run] Shutdown server: %s", err))
	}
}
