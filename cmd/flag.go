package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ontio/mdserver/md5config"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"strings"
)

var (
	LogLevelFlag = cli.UintFlag{
		Name:  "loglevel",
		Usage: "Set the log level to `<level>` (0~6). 0:Trace 1:Debug 2:Info 3:Warn 4:Error 5:Fatal 6:MaxLevel",
		Value: uint(md5config.DEFAULT_LOG_LEVEL),
	}
	ConfigfileFlag = cli.StringFlag{
		Name:   "config",
		Usage:  "specify configfile",
		Value:  "config.json",
		EnvVar: "CONFIG_FILE",
	}
)

func GetFlagName(flag cli.Flag) string {
	name := flag.GetName()
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.Split(name, ",")[0])
}

func GetServerConfig(ctx *cli.Context) (*md5config.Config, error) {
	cf := ctx.String(GetFlagName(ConfigfileFlag))
	file, err := os.Open(cf)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	cfg := &md5config.Config{}
	err = json.Unmarshal(bs, cfg)
	if err != nil {
		return nil, err
	}

	if cfg.LevelDBName == "" || cfg.RestPort == "" {
		return nil, fmt.Errorf("need config DB name or restPort")
	}

	return cfg, nil
}
