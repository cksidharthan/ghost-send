package router

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	validator "github.com/asaskevich/govalidator"
	"github.com/cksidharthan/share-secret/pkg/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var bootstrapOnly sync.Once

type Router struct {
	fx.Out

	Engine    *gin.Engine
	BaseRoute *gin.RouterGroup `name:"baseRoutes"`
	Server    *http.Server
}

// New - creates a new router instance and serves it to the application - using uber fx.
func New(lc fx.Lifecycle, envCfg *config.Config, zapLog *zap.SugaredLogger) Router {
	// these calls mutate a global state
	// using sync.Once here prevents data races prevents data races
	bootstrapOnly.Do(func() {
		gin.SetMode("release")

		// Setup validator to ignore fields with no `validate` tag
		validator.SetFieldsRequiredByDefault(false)
	})

	engine := gin.Default()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.RedirectTrailingSlash = false

	// Configure CORS
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	server := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      engine,
		Addr:         fmt.Sprintf(":%d", envCfg.Port),
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// Setup server
			errCh := make(chan error)

			go func(errCh chan error) {
				if err := server.ListenAndServe(); err != nil {
					errCh <- err
				}
			}(errCh)

			// Wait 1s for errors
			select {
			case err := <-errCh:
				return err
			case <-time.After(time.Second):
				return nil
			}
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return Router{
		Engine:    engine,
		BaseRoute: engine.Group(""),
		Server:    server,
	}
}
