package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"user-api/internal/config"
	"user-api/internal/db"
	"user-api/pkg/user"
	"user-api/pkg/web"

	"github.com/shipperizer/miniature-monkey/utils"
	"github.com/spf13/viper"
	//monitoringLibrary "github.com/some/monitoring/library"
)

func main() {
	config.LoadConfig()

	//monitor := monitoringLibrary.NewMonitor()
	logger := utils.NewLogger("info")

	db, _ := db.Init(viper.GetString("database.user"), viper.GetString("database.password"), viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.name"))

	api := web.NewApi(
		user.NewService(user.NewRepo(db)),
		logger,
	)

	fmt.Println("Here 1")

	srv := &http.Server{
		Addr: "0.0.0.0:8001",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      api.Handler(),
	}

	go func() {
		fmt.Println("Here 2")
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
		fmt.Println("Here 3")
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	fmt.Println("Here 4")
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logger.Info("Shutting down")
	os.Exit(0)

}
