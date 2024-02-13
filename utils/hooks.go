package utils

import (
	"fmt"
	"github.com/dimixlol/hosts-proxy/logging"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"github.com/loopfz/gadgeto/tonic"
	"net/http"
	"strings"
)

type SuccessfulResponse struct {
	Status int         `json:"status" example:"200"`
	Data   interface{} `json:"data"`
}

type UnsuccessfulResponse struct {
	Status int         `json:"status" example:"500"`
	Err    interface{} `json:"error"`
}

func NewSuccessfulResponse(status int, data interface{}) *SuccessfulResponse {
	return &SuccessfulResponse{
		Status: status,
		Data:   data,
	}
}

func NewUnsuccessfulResponse(status int, data interface{}) *UnsuccessfulResponse {
	return &UnsuccessfulResponse{
		Status: status,
		Err:    data,
	}
}

func NewSuccessfulResponseWithCode(status int, data interface{}) (int, *SuccessfulResponse) {
	return status, NewSuccessfulResponse(status, data)
}

func NewUnsuccessfulResponseWithCode(status int, err interface{}) (int, *UnsuccessfulResponse) {
	return status, NewUnsuccessfulResponse(status, err)
}

func Recovery(c *gin.Context, err any) {
	logger := logging.GetLogger(c)
	c.AbortWithStatusJSON(NewUnsuccessfulResponseWithCode(http.StatusInternalServerError, "internal server error"))
	logger.Errorf(c, "internal server error: %s", err)
}

func RenderHook(c *gin.Context, statusCode int, data interface{}) {
	var status int
	var fn func(code int, obj any)

	if c.Writer.Written() {
		status = c.Writer.Status()
	} else {
		status = statusCode
	}

	if gin.IsDebugging() {
		fn = c.IndentedJSON
	} else {
		fn = c.JSONP
	}

	if status >= http.StatusBadRequest {
		fn(NewUnsuccessfulResponseWithCode(status, data))
	} else {
		fn(NewSuccessfulResponseWithCode(status, data))
	}
}

func BindingHook(c *gin.Context, obj interface{}) error {
	err := tonic.DefaultBindingHook(c, obj)
	logger := logging.GetLogger(c)
	if err != nil {
		logger.Errorf(c, "binding error: %s", err.Error())
		return fmt.Errorf("incomplete request")
	}
	return nil
}

func ErrHook(c *gin.Context, e error) (int, interface{}) {
	errcode, errpl := http.StatusInternalServerError, e.Error()
	if _, ok := e.(tonic.BindError); ok {
		// Hide bind errors response
		errcode, errpl = http.StatusBadRequest, strings.Replace(e.Error(), "binding error: ", "", 1)
	} else {
		switch {
		case errors.Is(e, errors.BadRequest) || errors.Is(e, errors.NotValid) || errors.Is(e, errors.AlreadyExists) || errors.Is(e, errors.NotSupported) || errors.Is(e, errors.NotAssigned) || errors.Is(e, errors.NotProvisioned):
			errcode, errpl = http.StatusBadRequest, e.Error()
		case errors.Is(e, errors.Forbidden):
			errcode, errpl = http.StatusForbidden, e.Error()
		case errors.Is(e, errors.MethodNotAllowed):
			errcode, errpl = http.StatusMethodNotAllowed, e.Error()
		case errors.Is(e, errors.NotFound) || errors.Is(e, errors.UserNotFound):
			errcode, errpl = http.StatusNotFound, e.Error()
		case errors.Is(e, errors.Unauthorized):
			errcode, errpl = http.StatusUnauthorized, e.Error()
		case errors.Is(e, errors.NotImplemented):
			errcode, errpl = http.StatusNotImplemented, e.Error()
		}
	}
	return errcode, errpl
}
