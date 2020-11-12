package core

import (
	"github.com/ontio/mdserver/md5config"
	"github.com/ontio/ontology/core/store/leveldbstore"
)

type ServerRun struct {
	Store *leveldbstore.LevelDBStore
}

var DefServerRun *ServerRun

func NewServerRunTime(config *md5config.Config) (*ServerRun, error) {
	store, err := leveldbstore.NewLevelDBStore(config.LevelDBName)
	if err != nil {
		return nil, err
	}

	return &ServerRun{
		Store: store,
	}, nil
}
