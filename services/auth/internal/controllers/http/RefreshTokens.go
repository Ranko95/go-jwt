package http

import (
	"auth/internal/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

const REFRESH_TOKENS_ENDPOINT = "/refresh"

type RefreshTokensController struct {
	logger     *infrastructure.Logger
	httpServer *infrastructure.HTTPServer
}

func NewRefreshTokensController(
	logger *infrastructure.Logger,
	httpServer *infrastructure.HTTPServer,
) *RefreshTokensController {
	return &RefreshTokensController{
		logger:     logger,
		httpServer: httpServer,
	}
}

func (controller *RefreshTokensController) Bind() {
	controller.httpServer.Router().POST(REFRESH_TOKENS_ENDPOINT, controller.handle)
}

func (controller *RefreshTokensController) handle(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
