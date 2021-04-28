package identity

import (
	"context"
	"encoding/json"
	"strings"
	"sync"
	"time"

	gin "github.com/gin-gonic/gin"
	resty "github.com/go-resty/resty/v2"
	goutil "github.com/lanvige/goutils"
)

// RequestIDKey RequestIDKey
var RequestIDKey = "X-Request-ID"
var runOnce sync.Once
var restyClient *resty.Client

// NewRequest NewRequest
func NewRequest(ctx context.Context) *resty.Request {
	return Client().R().SetContext(ctx)
}

// Client Client
func Client() *resty.Client {
	runOnce.Do(func() {
		restyClient = resty.New().
			SetHeader("Content-Type", "application/json").
			SetHeader("Charset", "utf-8").
			SetTimeout(10 * time.Second)
	})

	return restyClient
}

// Execute Execute
func Execute(request *resty.Request, method, url string, body interface{}, resp interface{}) error {
	if body != nil {
		request = request.SetBody(body)
	}

	r, err := request.Execute(strings.ToUpper(method), url)
	if err != nil {
		return err
	}

	//检查requestid
	sourceReqID := request.Header.Get(RequestIDKey)
	returnReqID := r.Header().Get(RequestIDKey)

	if r.IsSuccess() && (sourceReqID == "" || returnReqID == "" || sourceReqID != returnReqID) {
		return NewAppError("RequestID Not Match")
	}

	return ParseResponse(r, resp)
}

// ParseResponse ParseResponse
func ParseResponse(r *resty.Response, obj interface{}) error {
	if r.IsSuccess() {
		if err := json.Unmarshal([]byte(r.Body()), obj); err != nil {
			return err
		}
		return nil
	}

	if r.IsError() {
		var appErr AppError
		if err := json.Unmarshal([]byte(r.Body()), &appErr); err != nil {
			return err
		}
		return &appErr
	}

	return nil
}

// GenRequestID GenRequestID
func GenRequestID(ctx context.Context) string {
	var requestID string

	if gin, ok := ctx.(*gin.Context); ok {
		reqID := gin.GetHeader(RequestIDKey)
		if reqID != "" {
			requestID = reqID
		}
	}

	if requestID == "" {
		if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
			if reqID != "" {
				requestID = reqID
			}
		}
	}

	if requestID == "" {
		requestID = goutil.UUIDV4StringGen()
	}

	return requestID
}
