package common

import (
	"github.com/NETkiddy/common-go/log"
	"net/http"
	"runtime/debug"
)

func ExceptionHandler(request *http.Request, ex interface{}) {
	log.LoggerFromContextWithCaller(request.Context()).Errorf("%+v", ex)
	log.LoggerFromContextWithCaller(request.Context()).Errorf("%+v", string(debug.Stack()))
}
