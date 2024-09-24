package controller

import (
	"embed"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/tang95/x-seek/config"
	"github.com/tang95/x-seek/internal/auth"
	"github.com/tang95/x-seek/internal/model"
	"github.com/tang95/x-seek/internal/service"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

//go:embed static
var consoleFS embed.FS

type Controller struct {
	service      *service.Service
	config       *config.Server
	logger       *zap.Logger
	transaction  service.Transaction
	auth         *auth.Auth
	incidentRepo model.IncidentRepo
	userRepo     model.UserRepo
	teamRepo     model.TeamRepo
}

func NewController(
	service *service.Service,
	config *config.Server,
	logger *zap.Logger,
	transaction service.Transaction,
	a *auth.Auth,
	incidentRepo model.IncidentRepo,
	userRepo model.UserRepo,
	teamRepo model.TeamRepo,
) (*Controller, error) {
	return &Controller{
		service:      service,
		config:       config,
		logger:       logger,
		transaction:  transaction,
		auth:         a,
		incidentRepo: incidentRepo,
		userRepo:     userRepo,
		teamRepo:     teamRepo,
	}, nil
}

func (controller *Controller) WithRoutes(engine *gin.Engine, jwtMiddleware *jwt.GinJWTMiddleware) {
	// console
	consoleServer := static.Serve("/", static.EmbedFolder(consoleFS, "static"))
	engine.Use(consoleServer)
	engine.NoRoute(func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodGet &&
			!strings.ContainsRune(ctx.Request.URL.Path, '.') &&
			!strings.HasPrefix(ctx.Request.URL.Path, "/api/") {
			ctx.Request.URL.Path = "/"
			consoleServer(ctx)
		}
	})
	// api group
	api := engine.Group("/api")

	// component
	component := api.Group("/incident", jwtMiddleware.MiddlewareFunc())
	component.GET("/query", controller.queryIncidents())
	// oauth
	oauth := api.Group("/auth")
	oauth.GET("/providers", controller.oauthProviders())
	oauth.POST("/validate", controller.oauthValidate(jwtMiddleware))
	oauth.GET("/authorizeUrl", controller.oauthAuthorizeUrl())
}
