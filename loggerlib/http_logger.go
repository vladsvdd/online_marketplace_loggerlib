// Package online_marketplace_libs logger/http_logger.go
package online_marketplace_loggerlib

import (
	"context"
	"fmt"
	"strings"
)

// HTTPRequestLog содержит информацию о HTTP-запросе для логирования
type HTTPRequestLog struct {
	Method      string                 `json:"method"`
	URL         string                 `json:"url"`
	Params      map[string]interface{} `json:"params,omitempty"`
	RequestData map[string]interface{} `json:"request_data,omitempty"`
	Response    interface{}            `json:"response,omitempty"`
	StatusCode  int                    `json:"status_code,omitempty"`
	Error       error                  `json:"error,omitempty"`
	Service     string                 `json:"service"`
}

// LogHTTPRequest логирует детали HTTP-запроса к внешнему API
func (s *Logger) LogHTTPRequest(ctx context.Context, req HTTPRequestLog) {
	attrs := []any{
		"method", req.Method,
		"url", req.URL,
		"service", req.Service,
	}

	if req.Params != nil && len(req.Params) > 0 {
		attrs = append(attrs, "params", req.Params)
	}

	if req.RequestData != nil && len(req.RequestData) > 0 {
		attrs = append(attrs, "request_data", req.RequestData)
	}

	if req.Response != nil {
		attrs = append(attrs, "response", req.Response)
	}

	if req.StatusCode > 0 {
		attrs = append(attrs, "status_code", req.StatusCode)
	}

	if req.Error != nil {
		attrs = append(attrs, "error", req.Error)
	}

	msg := fmt.Sprintf("[%s] API request details ", req.Service)
	s.With(attrs...).Info(ctx, msg)
}

// LogAPIRequestWithURL логирует детали HTTP-запроса с автоматическим построением URL
func (s *Logger) LogAPIRequestWithURL(ctx context.Context, serviceName, method, baseURL string, params map[string]interface{}, requestData map[string]interface{}, response interface{}, statusCode int, err error) {
	url := baseURL
	if params != nil && len(params) > 0 {
		var queryParams []string
		for key, value := range params {
			queryParams = append(queryParams, fmt.Sprintf("%s=%v", key, value))
		}
		if len(queryParams) > 0 {
			url += "?" + strings.Join(queryParams, "&")
		}
	}

	req := HTTPRequestLog{
		Method:      method,
		URL:         url,
		Params:      params,
		RequestData: requestData,
		Response:    response,
		StatusCode:  statusCode,
		Error:       err,
		Service:     serviceName,
	}

	s.LogHTTPRequest(ctx, req)
}
