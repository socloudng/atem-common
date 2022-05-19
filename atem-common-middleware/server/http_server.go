package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"go.uber.org/zap"
)

func NewHttpServer(servCfg *ServerConfig, httpHandler http.Handler, logger *zap.Logger) {
	addr := ":" + strconv.Itoa(servCfg.ServerPort)
	serv := &http.Server{
		Addr:           addr,
		Handler:        httpHandler,
		ReadTimeout:    time.Duration(servCfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(servCfg.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	logger.Info("\nserver run success on " + addr)
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := serv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout.
			logger.Error("HTTP server Shutdown: %v", zap.Error(err))
		}

		logger.Info("HTTP server Shutdown!")
		close(idleConnsClosed)
	}()

	if servCfg.HttpsEnabled {
		certPath, keyPath := servCfg.HttpsCertPath, servCfg.HttpsKeyPath
		if err := serv.ListenAndServeTLS(certPath, keyPath); err != http.ErrServerClosed {
			logger.Error("HTTPS server ListenAndServe: %v", zap.Error(err))
		}
	} else {
		if err := serv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Error("HTTP server ListenAndServe: %v", zap.Error(err))
		}
	}

	<-idleConnsClosed
}

// EnableCrossDomain sets the `Access-Control-Allow-Methods` header and the
// `Access-Control-Allow-Origin` header to the response to enable cross domain.
//
// TODO: We should restrict the origin, and may set in `config.toml`.
func EnableCrossDomain(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin == "" { // not cross origin
		return
	}

	header := w.Header()
	header.Set("Access-Control-Allow-Methods", "OPTIONS,POST,GET")
	header.Set("Access-Control-Allow-Origin", origin)
}
