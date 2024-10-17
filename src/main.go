package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gin "github.com/gin-gonic/gin"
	config "github.com/mzamani18/rapd_solutions_challenge/config"
	controller "github.com/mzamani18/rapd_solutions_challenge/controller"
	openaiservice "github.com/mzamani18/rapd_solutions_challenge/open_ai_services"
	"github.com/mzamani18/rapd_solutions_challenge/utils"
	log "github.com/sirupsen/logrus"
)


func init() {
	config.LoadConfig("config/config.json")
	openaiservice.InitilizeClient()
	utils.InitializeTrie()
}


func main() {
	log.SetReportCaller(true)
    router := gin.Default()

	router.POST("/v1/convert/", controller.ConvertDataToStructuredData)
	router.POST("/v1/convert/batch/", controller.ConvertBatchDataToStructuredData)
	router.GET("/v1/documents/", controller.GetAllDocuments)
	router.GET("/v1/search/", controller.SearchOnStructuredData)

    srv := &http.Server{
        Addr:    config.Config.Listen,
        Handler: router,
		ReadTimeout: 10 * time.Second,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    // Wait for interrupt signal to gracefully shutdown the server with a timeout.
    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }

    log.Println("Server exiting")
}
