package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sonata-labs/sonata/config"
	"github.com/sonata-labs/sonata/gen/api/v1/v1connect"
	"github.com/sonata-labs/sonata/types/module"
	"go.uber.org/zap"
)

type Server struct {
	*module.BaseModule
	config *config.Config

	httpServer *echo.Echo

	chain       v1connect.ChainHandler
	storage     v1connect.StorageHandler
	system      v1connect.SystemHandler
	p2p         v1connect.P2PHandler
	ddex        v1connect.DDEXHandler
	composition v1connect.CompositionHandler
	account     v1connect.AccountHandler
	validator   v1connect.ValidatorHandler
}

func (s *Server) Name() string {
	return "server"
}

var _ module.Module = (*Server)(nil)

func NewServer(config *config.Config, logger *zap.Logger, chain v1connect.ChainHandler, storage v1connect.StorageHandler, system v1connect.SystemHandler, p2p v1connect.P2PHandler, ddex v1connect.DDEXHandler, composition v1connect.CompositionHandler, account v1connect.AccountHandler, validator v1connect.ValidatorHandler) (*Server, error) {
	httpServer := echo.New()

	svc := &Server{
		config:      config,
		httpServer:  httpServer,
		chain:       chain,
		storage:     storage,
		system:      system,
		p2p:         p2p,
		ddex:        ddex,
		composition: composition,
		account:     account,
		validator:   validator,
	}
	svc.BaseModule = module.NewBaseModule(logger.Named(svc.Name()))
	return svc, nil
}

func (s *Server) Start() error {
	s.AwaitStartupDeps()
	s.Logger.Info("starting")

	s.registerRoutes()

	address := fmt.Sprintf("%s:%d", s.config.Sonata.HTTP.Host, s.config.Sonata.HTTP.Port)
	s.MarkReady()

	if err := s.httpServer.Start(address); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("starting http server: %w", err)
	}
	return nil
}

func (s *Server) Stop() error {
	s.AwaitShutdownDeps()
	s.Logger.Info("stopping")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	var err error
	if shutdownErr := s.httpServer.Shutdown(shutdownCtx); shutdownErr != nil && shutdownErr != http.ErrServerClosed {
		err = shutdownErr
	}
	s.MarkStopped()
	return err
}
