package Response

import (
	"database/sql"
	"encoding/json"
	"go-rest-api/Constant"
	"net/http"
)

func ResponseSuccess(w http.ResponseWriter, body interface{}, internalCode Constant.InternalCode) {
	result := ResponseSuccessStruct{
		RestCode:    int(internalCode),
		RestStatus:  internalCode.Response().HttpTitle,
		RestMessage: internalCode.Response().Description,
		RestResult:  body,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(internalCode.Response().HttpCode)
	json.NewEncoder(w).Encode(result)
}

func ResponseError(w http.ResponseWriter, err error, internalCode Constant.InternalCode) {
	result := ResponseErrorStruct{
		RestCode:    int(internalCode),
		RestStatus:  internalCode.Response().HttpTitle,
		RestMessage: internalCode.Response().Description,
	}

	if err != nil {
		result.RestResult = append(result.RestResult, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(internalCode.Response().HttpCode)
	json.NewEncoder(w).Encode(result)
}

func ResponseService(HasErr bool, Err error, InternalCode Constant.InternalCode, tx *sql.Tx, Result interface{}) (result RespResultService) {
	result = RespResultService{
		HasErr:       HasErr,
		Err:          Err,
		InternalCode: int(InternalCode),
		Tx:           tx,
		Result:       Result,
	}
	return
}
