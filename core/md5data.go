package core

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/store/leveldbstore"
)

const (
	TYPE_DIANXING uint8 = 1
	TYPE_LIANTONG uint8 = 2
	TYPE_YIDONG   uint8 = 3
)

type PhoneMD5 struct {
	PType       uint8          `json:"pType"`
	PhoneNumber uint64         `json:"phoneNumber"`
	PhoneMD5    [md5.Size]byte `json:"phoneMD5"`
}

func (self *PhoneMD5) MarshalJSON() ([]byte, error) {
	type t struct {
		PType       uint8  `json:"pType"`
		PhoneNumber uint64 `json:"phoneNumber"`
		PhoneMD5    string `json:"phoneMD5"`
	}

	x := &t{
		PType:       self.PType,
		PhoneNumber: self.PhoneNumber,
		PhoneMD5:    fmt.Sprintf("%x", self.PhoneMD5[:]),
	}

	return json.Marshal(x)
}

func BatchPutPhoneMD5(info *PhoneMD5, store *leveldbstore.LevelDBStore) {
	sink := common.NewZeroCopySink(nil)
	sink.WriteUint8(info.PType)
	sink.WriteUint64(info.PhoneNumber)
	sink.WriteBytes(info.PhoneMD5[:])

	store.BatchPut(info.PhoneMD5[:], sink.Bytes())
}

func GetPhoneMD5(md5Key []byte) (*PhoneMD5, error) {
	log.Debugf("GetPhoneMD5. Y.0 %x", md5Key)
	bytes, err := DefServerRun.Store.Get(md5Key)
	if err != nil {
		return nil, err
	}

	source := common.NewZeroCopySource(bytes)
	pType, eof := source.NextUint8()
	if eof {
		return nil, fmt.Errorf("GetPhoneMD5.N.0 eof")
	}

	phoneNumber, eof := source.NextUint64()
	if eof {
		return nil, fmt.Errorf("GetPhoneMD5.N.1 eof")
	}

	md5Data, eof := source.NextBytes(md5.Size)
	if eof {
		return nil, fmt.Errorf("GetPhoneMD5.N.1 eof")
	}

	data := &PhoneMD5{
		PType:       pType,
		PhoneNumber: phoneNumber,
	}

	copy(data.PhoneMD5[:], md5Data)

	return data, nil
}
