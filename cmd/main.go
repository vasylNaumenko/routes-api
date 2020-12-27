/*
 * telegram: @VasylNaumenko
 */

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"routes-api/pkg/api/v1/rest"
	"routes-api/pkg/api/v1/rest/handlers"
	"routes-api/pkg/api/v1/services/osrm"
	"routes-api/pkg/config"
	"routes-api/pkg/initialize"
)

const (
	serviceName = "routes-api"
)

func main() {
	// Init config
	cfg, err := config.Init(config.MainConfigFile)
	if err != nil {
		fmt.Printf("failed to init cfg: %s", err)
		os.Exit(1)
	}

	// Init logger
	logger, err := initialize.Logger(&cfg.Logger, serviceName)
	if err != nil {
		fmt.Printf("failed to init log: %s", err)
		os.Exit(1)
	}

	logger.Infof("Starting %s service", serviceName)

	// init API handlers and services
	handlerMeta := handlers.NewMeta(logger) // health check

	osrmService := osrm.New(cfg.RestVendors, logger)
	handlerOsrm := handlers.NewOsrm(logger, osrmService)

	// Init REST API
	api := rest.New(cfg.APIServer, logger, handlerMeta, handlerOsrm)
	api.Run()

	// Listening to service shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	s := <-c
	logger.Infof("received signal %q", s)
	os.Exit(0)
}
