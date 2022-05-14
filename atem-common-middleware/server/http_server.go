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
