package md5config

import (
	"github.com/ontio/ontology/common/log"
)

var Version = ""

var (
	DEFAULT_LOG_LEVEL uint = log.InfoLog
)

type Config struct {
	RestPort    string `json:"restPort"`
	LevelDBName string `json:"levelDBName"`
}
