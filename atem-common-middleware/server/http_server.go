package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/socloudng/atem-common/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewHttpServer(servCfg *configs.ServerConfig, httpHandler http.Handler, logger *zap.Logger) {
	addr := strconv.Itoa(servCfg.ServerPort)
	serv := &http.Server{
		Addr:           addr,
		Handler:        httpHandler,
		ReadTimeout:    time.Duration(servCfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(servCfg.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	logger.Info("server run success on ", zapcore.Field{Key: "address", String: addr})
	fmt.Println("Welcom to " + servCfg.ServerName)
	logger.Error(serv.ListenAndServe().Error())
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
