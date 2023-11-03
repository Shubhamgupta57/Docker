package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"devops-testing/api"
	"devops-testing/api/middleware"
	config "devops-testing/cmd/config"
	"devops-testing/logger"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

// Server := server object
type Server struct {
	API        *api.API
	httpServer *http.Server
	Router     *mux.Router
	Logger     *logger.Logger
	Config     *config.Config
}

// NewServer := returns new server instance
func NewServer() (*Server, error) {
	config := config.GetConfig()
	logger := logger.NewLogger(true, true, "server")
	router := mux.NewRouter()
	server := Server{
		Router: router,
		Config: config,
		Logger: logger,
	}
	return &server, nil
}

// StartServer := run server on a port
func (s *Server) StartServer() {
	n := negroni.New()
	c := cors.New(cors.Options{
		AllowedOrigins:   s.Config.CORSConfig.AllowedOrigins,
		AllowedMethods:   s.Config.CORSConfig.AllowedMethods,
		AllowCredentials: s.Config.CORSConfig.AllowCredentials,
		AllowedHeaders:   s.Config.CORSConfig.AllowedHeaders,
	})
	n.Use(c)

	s.API = api.NewAPI(s.Router, s.Config)
	requestLogger := logger.NewLogger(true, true, "requests")
	// csrfMiddleware := csrf.Protect([]byte(s.Config.ServerConfig.CSRFProtect))
	n.UseFunc(middleware.NewRequestLoggerMiddleware(requestLogger).GetMiddlewareHandler())
	// n.UseFunc(middleware.NewAuthenticationMiddleware(s.API.Session).GetMiddlewareHandler())
	// s.Router.Use(csrfMiddleware)
	n.UseHandler(s.Router)

	listenArr := fmt.Sprintf("%s:%s", s.Config.ServerConfig.Addr, s.Config.ServerConfig.Port)
	s.httpServer = &http.Server{
		Handler: n,
		Addr:    listenArr,
	}
	s.Logger.Log.Debug().Msgf("Starting Server at %s", listenArr)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil {
			s.Logger.Log.Err(err)
			return
		}
	}()
	s.Logger.Log.Debug().Msg("Server Running")
	wg.Wait()
}

// StopServer := stops running server
func (s *Server) StopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.httpServer.Shutdown(ctx)
	os.Exit(0)
}
