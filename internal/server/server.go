package server

import (
	_ "github.com/estromenko/yoof-api/docs"
	services "github.com/estromenko/yoof-api/internal/services"
	"github.com/estromenko/yoof-api/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	Host string `json:"host" yaml:"host" mapstructure:"host"`
	Port string `json:"port" yaml:"port" mapstructure:"port"`

	Services *services.Config `mapstructure:"services"`
}

type Server struct {
	logger *zerolog.Logger
	config *Config
	db     *db.Database

	services *services.Services
}

func New(db *db.Database, logger *zerolog.Logger, config *Config) *Server {
	return &Server{
		logger: logger,
		db:     db,
		config: config,
	}
}

func (s *Server) Services() *services.Services {
	return s.services
}

func (s *Server) Run() error {
	s.services = services.InitServices(s.db.DB(), s.config.Services)

	gin.SetMode(gin.ReleaseMode)
	r := s.route()

	s.logger.Debug().Msg("Server started at " + s.config.Host + ":" + s.config.Port)
	return r.Run(s.config.Host + ":" + s.config.Port)
}

// @Description Base page
// @Accept  json
// @Produce  json
// @Router / [get]

func (s *Server) controller(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{
		"ok": "ok",
	})
}

func (s *Server) route() *gin.Engine {
	r := gin.New()

	public := r.Group("/")
	{
		public.Use(s.baseMiddleware())

		auth := public.Group("/auth")
		{
			auth.POST("/reg", s.RegistrationHandler())
			auth.POST("/login", s.LoginHandler())
		}

		user := public.Group("/user")
		{
			user.Use(s.authMiddleware())

			user.GET("/info", s.GetUserInfoHandler())
		}
	}

	private := r.Group("/")
	private.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://"+s.config.Host+":"+s.config.Port+"/docs/doc.json")))
	private.Any("/", s.controller)
	return r
}
