package http

import (
	"auth/internal/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

const LOGIN_ENDPOINT = "/login/:id"

type LoginController struct {
	logger     *infrastructure.Logger
	httpServer *infrastructure.HTTPServer
}

func NewLoginController(
	logger *infrastructure.Logger,
	httpServer *infrastructure.HTTPServer,
) *LoginController {
	return &LoginController{
		logger:     logger,
		httpServer: httpServer,
	}
}

func (controller *LoginController) Bind() {
	controller.httpServer.Router().POST(LOGIN_ENDPOINT, controller.handle)
}

func (controller *LoginController) handle(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
