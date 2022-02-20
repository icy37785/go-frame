package main

import (
	"fmt"
	"github.com/icy37785/go-frame/app/services/test-api/handlers"
	"github.com/icy37785/go-frame/pkg/app"
	"github.com/icy37785/go-frame/pkg/config"
	"github.com/icy37785/go-frame/pkg/logger"
	"github.com/icy37785/go-frame/pkg/storage/orm"
	"github.com/icy37785/go-frame/pkg/storage/sql"
	"github.com/icy37785/go-frame/pkg/util"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"os"
	"runtime"
)

var build = "develop"

func main() {

	// Construct the application logger.
	log, err := logger.New("TEST-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {

	// =========================================================================
	// GOMAXPROCS

	// Want to see what maxprocs reports.
	opt := maxprocs.Logger(log.Infof)

	// Set the correct number of threads for the service
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(opt); err != nil {
		return fmt.Errorf("maxprocs: %w", err)
	}
	log.Infow("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// =========================================================================
	// Configuration
	c := config.New("config")
	var cfg app.Config
	if err := c.Load("app", &cfg); err != nil {
		panic(err)
	}

	// =========================================================================
	// App Starting

	log.Infow("starting service", "version", build)
	defer log.Infow("shutdown complete")

	app.Conf = &cfg

	// =========================================================================
	// Initialize authentication support

	// =========================================================================
	// Database Support

	var ormConfig orm.Config
	var sqlConfig sql.Config
	if err := c.Load("database", &ormConfig); err != nil {
		panic(err)
	}
	_ = util.Copy(&ormConfig, &sqlConfig)
	ormDB := orm.NewOrm(&ormConfig)
	sqlDB := sql.NewSql(&sqlConfig)
	defer func() {
		log.Infow("shutdown", "status", "stopping database support", "host", ormConfig.Addr)
		db, _ := ormDB.DB()
		_ = db.Close()
		_ = sqlDB.Close()
	}()

	// =========================================================================
	// Start API Service

	log.Infow("startup", "status", "initializing V1 API support")

	appHTTP := handlers.NewHTTPServer(handlers.Config{
		Log: log,
		Orm: ormDB,
		DB:  sqlDB,
	})

	ac := app.AppsConfig{
		App: appHTTP,
		Log: log,
		Cfg: &cfg,
	}
	// Start server (with or without graceful shutdown).
	if build == "build" {
		app.StartServer(ac)
	} else {
		app.StartServerWithGracefulShutdown(ac)
	}

	return nil
}
