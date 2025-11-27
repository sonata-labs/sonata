package app

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sonata-labs/sonata/config"
	"github.com/sonata-labs/sonata/gen/api/v1/v1connect"
	"github.com/sonata-labs/sonata/store/chainstore"
	pebble_chainstore "github.com/sonata-labs/sonata/store/chainstore/pebble"
	"github.com/sonata-labs/sonata/store/localstore"
	pebble_localstore "github.com/sonata-labs/sonata/store/localstore/pebble"
	"github.com/sonata-labs/sonata/x/account"
	"github.com/sonata-labs/sonata/x/chain"
	"github.com/sonata-labs/sonata/x/composition"
	"github.com/sonata-labs/sonata/x/ddex"
	"github.com/sonata-labs/sonata/x/p2p"
	"github.com/sonata-labs/sonata/x/storage"
	"github.com/sonata-labs/sonata/x/system"
	"github.com/sonata-labs/sonata/x/validator"
	"golang.org/x/sync/errgroup"
)

type App struct {
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

	chainStore chainstore.ChainStore
	localStore localstore.LocalStore
}

func NewApp(config *config.Config) (*App, error) {
	chain := chain.NewChainService(config)
	storage := storage.NewStorageService(config)
	system := system.NewSystemService(config)
	p2p := p2p.NewP2PService(config)
	ddex := ddex.NewDDEXService(config)
	composition := composition.NewCompositionService(config)
	account := account.NewAccountService(config)
	validator := validator.NewValidatorService(config)

	chainStore, err := pebble_chainstore.NewPebbleChainStore(config.ChainStore.Path)
	if err != nil {
		return nil, err
	}

	localStore, err := pebble_localstore.NewPebbleLocalStore(config.LocalStore.Path)
	if err != nil {
		return nil, err
	}

	httpServer := echo.New()

	httpServer.HideBanner = true

	httpServer.Use(middleware.Logger(), middleware.Recover())

	httpServer.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]uint{"a": 440})
	})

	rpcGroup := httpServer.Group("")
	chainPath, chainHandler := v1connect.NewChainHandler(chain)
	rpcGroup.Any(chainPath, echo.WrapHandler(chainHandler))

	storagePath, storageHandler := v1connect.NewStorageHandler(storage)
	rpcGroup.Any(storagePath, echo.WrapHandler(storageHandler))

	systemPath, systemHandler := v1connect.NewSystemHandler(system)
	rpcGroup.Any(systemPath, echo.WrapHandler(systemHandler))

	p2pPath, p2pHandler := v1connect.NewP2PHandler(p2p)
	rpcGroup.Any(p2pPath, echo.WrapHandler(p2pHandler))

	ddexPath, ddexHandler := v1connect.NewDDEXHandler(ddex)
	rpcGroup.Any(ddexPath, echo.WrapHandler(ddexHandler))

	compositionPath, compositionHandler := v1connect.NewCompositionHandler(composition)
	rpcGroup.Any(compositionPath, echo.WrapHandler(compositionHandler))

	accountPath, accountHandler := v1connect.NewAccountHandler(account)
	rpcGroup.Any(accountPath, echo.WrapHandler(accountHandler))

	validatorPath, validatorHandler := v1connect.NewValidatorHandler(validator)
	rpcGroup.Any(validatorPath, echo.WrapHandler(validatorHandler))

	return &App{
		config:      config,
		httpServer:  httpServer,
		storage:     storage,
		system:      system,
		p2p:         p2p,
		ddex:        ddex,
		composition: composition,
		account:     account,
		validator:   validator,
		chainStore:  chainStore,
		localStore:  localStore,
	}, nil
}

func (app *App) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	defer func() {
		log.Printf("shutdown complete\n")
	}()

	eg.Go(func() error {
		return app.runHTTP(ctx)
	})

	eg.Go(func() error {
		<-ctx.Done()
		log.Printf("shutting down...\n")
		return app.Shutdown()
	})

	return eg.Wait()
}

func (app *App) Shutdown() error {
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	wg := sync.WaitGroup{}

	wg.Go(func() {
		if err := app.httpServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("failed to shutdown HTTP server: %v", err)
		}
	})

	wg.Wait()

	return nil
}
