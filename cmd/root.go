package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zucchinidev/go-file-store/infrastructure/logger"
	"github.com/zucchinidev/go-file-store/infrastructure/www"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var rootCmd = &cobra.Command{
	Use:   "file-storage",
	Short: "file-storage is a very fast file storage with different buckets implemented",
	Run: func(cmd *cobra.Command, args []string) {
		shutdown := make(chan struct{}, 1)
		svc := www.Server(www.Conf{Addr: "0.0.0.0:3000"})
		kill := terminate(shutdown, svc)
		l := logger.New()
		go func(l *logger.Standard, svc *http.Server) {
			if err := svc.ListenAndServe(); err != http.ErrServerClosed {
				l.HTTPServerError(err)
				kill(l)
			}
		}(l, svc)
		go interruptSignal(l, shutdown)
		<-shutdown
	},
}

// Execute assembles the app commands necessaries to up the applications
func Execute() {
	rootCmd.AddCommand(liveCommand)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func terminate(shutdown chan struct{}, svc *http.Server) func(l *logger.Standard) {
	return func(l *logger.Standard) {
		if err := svc.Shutdown(context.Background()); err != nil {
			l.HTTPServerShutdownError(err)
		}
		shutdown <- struct{}{}
	}
}

func interruptSignal(l *logger.Standard, shutdown chan struct{}) {
	signals := make(chan os.Signal, 1)
	// sigterm signal sent from kubernetes, interrupt signal sent from terminal
	signal.Notify(signals, syscall.SIGTERM, os.Interrupt)
	<-signals
	l.ReceivedInterruptSignal()
	shutdown <- struct{}{}
}
