package Response

import "database/sql"

type ResponseSuccessStruct struct {
	RestCode    int         `json:"code"`
	RestStatus  string      `json:"status"`
	RestMessage string      `json:"message"`
	RestResult  interface{} `json:"result" swaggertype:"object,string" example:"key:value,key2:value2"`
}

type ResponseErrorStruct struct {
	RestCode    int           `json:"code"`
	RestStatus  string        `json:"status"`
	RestMessage string        `json:"message"`
	RestResult  []interface{} `json:"result" swaggertype:"object,string" example:"key:value,key2:value2"`
}

type RespResultService struct {
	HasErr       bool
	Err          error
	InternalCode int
	Tx           *sql.Tx
	Result       interface{}
}
