package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	goutil "github.com/lanvige/goutils"
)

var requestIDKey = "X-Request-ID"
var runOnce sync.Once
var restyClient *resty.Client
var RequestIDKey = "X-Request-ID"

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
	fmt.Printf("url:%s", url)

	if body != nil {
		request = request.SetBody(body)
	}

	fmt.Printf("request:%v", request)
	r, err := request.Execute(strings.ToUpper(method), url)
	if err != nil {
		return err
	}

	// 检查requestid
	sourceReqID := request.Header.Get(requestIDKey)
	returnReqID := r.Header().Get(requestIDKey)
	if sourceReqID == "" || returnReqID == "" || sourceReqID != returnReqID {
		return errors.New("RequestID Not Match")
	}

	fmt.Printf("resp.status:%s", r.Status())

	return ParseResponse(r, resp)
}

// ParseResponse ParseResponse
func ParseResponse(r *resty.Response, obj interface{}) error {
	if r.IsSuccess() {
		if obj != nil {
			e := json.Unmarshal(r.Body(), obj)
			if e != nil {
				fmt.Printf("parseResponse:%s", e.Error())
				return e
			}
			return nil
		}

		return nil
	}

	return fmt.Errorf("%s", r.Status())
}

// GenRequestID GenRequestID
func GenRequestID(ctx context.Context) string {
	var requestID string
	if gin, ok := ctx.(*gin.Context); ok {
		reqID := gin.GetHeader(requestIDKey)
		if reqID != "" {
			requestID = reqID
		}
	}

	if requestID == "" {
		if reqID, ok := ctx.Value(requestIDKey).(string); ok {
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