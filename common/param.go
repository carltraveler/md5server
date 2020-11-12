package common

const (
	SUCCESS      = 0
	NOTFOUND     = 1
	MD5LENERROR  = 2
	MD5DATAERROR = 3
)

var CodeMessageMap = map[int32]string{
	SUCCESS:      "success",
	NOTFOUND:     "md5 data not found",
	MD5LENERROR:  "md5 len should be 32",
	MD5DATAERROR: "md5 decode to hex error",
}

type MD5Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}
