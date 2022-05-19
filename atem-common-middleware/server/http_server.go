package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type httpServer struct {
	servCfg         *ServerConfig
	serv            *http.Server
	logger          *zap.Logger
	idleConnsClosed chan os.Signal
}

func NewHttpServer(servCfg *ServerConfig, httpHandler http.Handler, logger *zap.Logger) *httpServer {
	addr := ":" + strconv.Itoa(servCfg.ServerPort)
	newServer := &httpServer{
		logger:  logger,
		servCfg: servCfg,

		serv: &http.Server{
			Addr:           addr,
			Handler:        httpHandler,
			ReadTimeout:    time.Duration(servCfg.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(servCfg.WriteTimeout) * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	logger.Info("server run success on ", zap.String("addr", addr))
	return newServer
}

func (h *httpServer) Start(delayTime int64, useWatch bool) *httpServer {
	if delayTime > 0 {
		go h.autoStop(delayTime) //自动延迟关闭
	}
	if useWatch { //启用优雅退出
		go h.startServer()
		h.watch(h.stopServer)
	} else {
		h.startServer()
	}
	return h
}

// 监听信号
func (h *httpServer) watch(callBack ...func() error) {
	h.idleConnsClosed = make(chan os.Signal, 1)
	signal.Notify(h.idleConnsClosed, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// wait for system signal
	s := <-h.idleConnsClosed
	close(h.idleConnsClosed)
	h.logger.Info("receive exit signal, run quit funcs", zap.String("signal", s.String()))
	for i := range callBack {
		if err := callBack[i](); err != nil {
			h.logger.Fatal("HTTP server ListenAndServe Fatal: %v", zap.Error(err))
		}
	}
	h.logger.Info("quit funcs exec over!")
}

func (h *httpServer) Stop() {
	if h.idleConnsClosed != nil {
		h.idleConnsClosed <- syscall.SIGQUIT
	} else {
		h.stopServer()
	}
}

func (h *httpServer) autoStop(delayTime int64) {
	time.Sleep(time.Duration(delayTime) * time.Second)
	h.Stop()
}

func (h *httpServer) startServer() {
	if h.servCfg.HttpsEnabled {
		certPath, keyPath := h.servCfg.HttpsCertPath, h.servCfg.HttpsKeyPath
		if err := h.serv.ListenAndServeTLS(certPath, keyPath); err != http.ErrServerClosed {
			h.logger.Fatal("HTTPS server ListenAndServeTLS: %v", zap.Error(err))
			os.Exit(0)
		}
	} else {
		if err := h.serv.ListenAndServe(); err != http.ErrServerClosed {
			h.logger.Fatal("HTTP server ListenAndServe: %v", zap.Error(err))
			os.Exit(0)
		}
	}
}

func (h *httpServer) stopServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return h.serv.Shutdown(ctx)
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
