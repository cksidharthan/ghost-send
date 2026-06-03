package ui

import (
	"github.com/cksidharthan/ghost-send/frontend"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Deps struct {
	fx.In

	Engine *gin.Engine
	Logger *zap.SugaredLogger
}

// Register sets up the embedded Nuxt SPA as the catch-all handler on the gin
// engine. API routes registered before this call take priority; everything else
// falls through to the SPA which handles client-side routing via index.html.
func Register(deps Deps) {
	handler, err := NewHandler(frontend.FS)
	if err != nil {
		deps.Logger.Errorw("failed to create UI handler", "error", err)
		return
	}
	deps.Engine.NoRoute(gin.WrapH(handler))
}
