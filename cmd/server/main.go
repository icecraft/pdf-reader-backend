package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/api"
	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/config"
	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/dao"
	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/log"
	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	configFile := flag.String("config_file", "/config/config.json", "config file used when program startup")
	port := flag.String("port", "8080", "web server port that used")
	logLevel := flag.Int("log_level", 5, "log level")
	encoding := flag.String("log_encoding", "console", "log encoding [console, json]")
	flag.Parse()

	var conf config.Config
	config.LoadConfigFromFile(*configFile, &conf)
	if conf.Debug {
		fmt.Printf("[debug] config: %v\n", conf)
		log.SetLogger(log.Development(int8(*logLevel), *encoding))
	} else {
		log.SetLogger(log.Production(int8(*logLevel), *encoding))
	}

	// do some init
	if err := dao.InitDAO(conf.Mysql, conf.TplFolder); err != nil {
		log.Error(err, "failed to init dao")
		os.Exit(1)
	}
	if err := dao.Sync2(); err != nil {
		log.Error(err, "failed to sync talbe")
		os.Exit(1)
	}
	// es init
	err := utils.InitEs(conf.ES)
	if err != nil {
		log.Error(err, "failed to init es")
		os.Exit(1)
	}

	// register route
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Cookie"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Cookie"},
		MaxAge:           12 * time.Hour,
	}
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	v1Resource := r.Group("")

	server, err := api.NewServer(&conf, *encoding)
	if err != nil {
		panic(err)
	}
	api.RegisterApi(conf.Router, conf.Host, v1Resource, server)

	// start server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", *port),
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(err, "failed to listen server")
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	for {
		s := <-quit

		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("Shutdown Server....")
			time.Sleep(10 * time.Second)
			if err := srv.Shutdown(context.TODO()); err != nil {
				log.Error(err, "shutdown server now")
			}
			log.Info("Server exited")
			os.Exit(0)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}