package app

import (
	"os"
	"os/signal"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(ac AppsConfig) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := ac.App.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			ac.Log.Infof("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := ac.App.Listen(ac.Cfg.HTTP.Addr); err != nil {
		ac.Log.Infof("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(ac AppsConfig) {
	// Run server.
	if err := ac.App.Listen(ac.Cfg.HTTP.Addr); err != nil {
		ac.Log.Infof("Oops... Server is not running! Reason: %v", err)
	}
}
