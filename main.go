package main

import (
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/ontio/mdserver/cmd"
	"github.com/ontio/mdserver/md5config"
	"github.com/ontio/mdserver/core"
	"github.com/ontio/mdserver/restful"
	"github.com/ontio/ontology/common/log"
)

func setupAPP() *cli.App {
	app := cli.NewApp()
	app.Usage = "Md5 server CLI"
	app.Action = startMd5Server
	app.Version = md5config.Version
	app.Copyright = "Copyright in 2018 The Ontology Authors"
	app.Flags = []cli.Flag{
		cmd.LogLevelFlag,
		cmd.ConfigfileFlag,
	}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func main() {
	if err := setupAPP().Run(os.Args); err != nil {
		log.Errorf("%s", err)
		os.Exit(1)
	}
}

func startMd5Server(ctx *cli.Context) {
	initLog(ctx)
	cfg, err := cmd.GetServerConfig(ctx)
	if err != nil {
		log.Errorf("startMd5Server.N.0 %s", err)
		return
	}

	serverRun, err := core.NewServerRunTime(cfg)
	if err != nil {
		log.Errorf("startMd5Server.N.1 %s", err)
		return
	}
	core.DefServerRun = serverRun

	restful.NewRouter()
	startServer(cfg)
	waitToExit()
}

func initLog(ctx *cli.Context) {
	logLevel := ctx.GlobalInt(cmd.GetFlagName(cmd.LogLevelFlag))
	log.InitLog(logLevel, log.Stdout)
}

func startServer(config *md5config.Config) {
	router := restful.NewRouter()
	go router.Run(":" + config.RestPort)
}

func waitToExit() {
	exit := make(chan bool, 0)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		for sig := range sc {
			log.Infof("mdserver server received exit signal: %s.", sig.String())
			close(exit)
			break
		}
	}()
	<-exit
}
