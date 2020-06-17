package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-api/config"
	"go-api/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	start := time.Now()
	gin.SetMode(config.AppMode)
	engine := gin.New()

	// 性能分析 - 正式环境不要使用！！！
	//pprof.Register(engine)

	// init config
	configFilePath := "config/conf/conf.yaml";
	if err := config.Init(configFilePath); err != nil {
		panic(err)
	}

	dbtype := viper.GetString("db.type")
	dbhostname:=viper.GetString("db.hostname")
	dbhostport:=viper.GetString("db.hostport")
	dbusername:=viper.GetString("db.username")
	dbpassword:=viper.GetString("db.password")
	dbdatabase:=viper.GetString("db.database")
	// init db
	config.InitDb(dbtype,dbhostname,dbhostport,dbusername,dbpassword,dbdatabase)

	// init redis
	//redis.Init()
	// init cache
	//cache.Init()

	// 设置路由
	route.InitRouter(engine)

	server := &http.Server{
		Addr:         config.AppPort,
		Handler:      engine,
		ReadTimeout:  config.AppReadTimeout * time.Second,
		WriteTimeout: config.AppWriteTimeout * time.Second,
	}

	fmt.Println("|-----------------------------------|")
	fmt.Println("|              go-api               |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|    Port" + config.AppPort + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()
	end := time.Since(start)
	fmt.Printf("运行时间：%v",end)

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
