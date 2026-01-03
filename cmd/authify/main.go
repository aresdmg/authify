package authify

import (
	"authify/internal/database"
	"authify/internal/middleware"
	"authify/internal/routes"
	"authify/pkg/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Start() {
	database.ConnectDB()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	routes.RouterEngine(r)

	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	server := &http.Server{
		Addr:    ":"+PORT,
		Handler: r,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.White(fmt.Sprintf("SERVER: %s", PORT))
		if err := server.ListenAndServe(); err != nil {
			logger.Red("SERVER FAILED TO START")
		}
	}()

	<-stop
	fmt.Println('\n')
	logger.White("SERVER stopped")

	defer func() {
		database.DisconnectDB()
	}()
}
