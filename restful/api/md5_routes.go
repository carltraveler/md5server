package api

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ontio/mdserver/common"
	"github.com/ontio/mdserver/core"
	"github.com/ontio/ontology/common/log"
	"net/http"
)

func RoutesApi(parent *gin.Engine) {
	apiRoute := parent.Group("/api")
	apiRoute.GET("/getMd5PhoneData/:md5", getMd5PhoneData)
}

func getMd5PhoneData(c *gin.Context) {
	md5 := c.Param("md5")

	if len(md5) != 32 {
		c.JSON(http.StatusOK, common.MD5Response{
			Code:    common.MD5LENERROR,
			Message: common.CodeMessageMap[common.MD5LENERROR],
		})
		return
	}

	log.Debugf("getMd5PhoneData: Y.0 %s", md5)

	md5Data, err := hex.DecodeString(md5)
	if err != nil {
		log.Debugf("getMd5PhoneData: N.0 %s", err)
		c.JSON(http.StatusOK, common.MD5Response{
			Code:    common.MD5DATAERROR,
			Message: common.CodeMessageMap[common.MD5DATAERROR] + fmt.Sprintf("%v", err),
		})
		return
	}

	result, err := core.GetPhoneMD5(md5Data)
	if err != nil {
		log.Debugf("getMd5PhoneData: N.1 %s", err)
		c.JSON(http.StatusOK, common.MD5Response{
			Code:    common.MD5DATAERROR,
			Message: common.CodeMessageMap[common.MD5DATAERROR] + fmt.Sprintf(" %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, common.MD5Response{
		Code:    common.SUCCESS,
		Message: common.CodeMessageMap[common.SUCCESS],
		Result:  result,
	})
}
