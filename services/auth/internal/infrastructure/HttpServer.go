package infrastructure

import (
	"auth/config"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	config config.HTTPServerConfig
	logger *Logger
	router *gin.Engine
}

func NewHttpServer(config config.HTTPServerConfig, logger *Logger) *HTTPServer {
	return &HTTPServer{
		config: config,
		logger: logger,
		router: gin.Default(),
	}
}

func (httpServer *HTTPServer) Router() *gin.Engine {
	return httpServer.router
}

func (httpServer *HTTPServer) Run() {
	httpServer.router.Run(httpServer.config.Address)
}
