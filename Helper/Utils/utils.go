package Utils

import (
	"context"
	"fmt"
	"go-rest-api/Constant"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type ContextMap struct {
	Id       string
	Name     string
	Username string
	Email    string
	Active   string
}

func SetValueContext(r *http.Request) (result ContextMap, contextMap jwt.MapClaims, err error) {

	ctxValue := r.Context().Value("claims_value")
	if ctxValue == nil {
		err = fmt.Errorf("claims_value not found in context")
		return
	}

	contextMap, ok := ctxValue.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("claims_value is not of type map[string]interface{}")
		return
	}

	result.Id = fmt.Sprintf("%v", contextMap["id"])
	result.Name = fmt.Sprintf("%v", contextMap["name"])
	result.Username = fmt.Sprintf("%v", contextMap["username"])
	result.Email = fmt.Sprintf("%v", contextMap["email"])
	result.Active = fmt.Sprintf("%v", contextMap["active"])
	return
}

func GetValueOfContext(key string, ctx context.Context) interface{} {
	if ctx.Value(key) != nil {
		return ctx.Value(key)
	}
	return ""
}

func TimeNow() time.Time {
	location, _ := time.LoadLocation(Constant.TimeLocation)
	return time.Now().In(location)
}
