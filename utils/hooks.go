package utils

import (
	"fmt"
	"github.com/dimixlol/knowyourwebsite/logging"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"github.com/loopfz/gadgeto/tonic"
	"net/http"
	"strings"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Err    interface{} `json:"error,omitempty"`
}

func (dr *Response) GetData() interface{} {
	return dr.Data
}
func (dr *Response) Error() string {
	return dr.Err.(string)
}
func (dr *Response) GetStatus() int {
	return dr.Status
}

func NewResponse(status int, data interface{}, err interface{}) *Response {
	return &Response{
		Status: status,
		Data:   data,
		Err:    err,
	}
}
func NewResponseWithCode(status int, data interface{}, err interface{}) (int, *Response) {
	return status, NewResponse(status, data, err)
}

func Recovery(c *gin.Context, err any) {
	logger := logging.GetLogger(c)
	c.AbortWithStatusJSON(NewResponseWithCode(http.StatusInternalServerError, nil, "internal server error"))
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
		fn(NewResponseWithCode(status, nil, data))
	} else {
		fn(NewResponseWithCode(status, data, nil))
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
