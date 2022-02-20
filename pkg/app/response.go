package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/icy37785/go-frame/pkg/errcode"
)

var resp *Response

func init() {
	resp = NewResponse()
}

// Response define a response struct
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Details []string    `json:"details,omitempty"`
}

// NewResponse return a response
func NewResponse() *Response {
	return &Response{}
}

// Success return a success response
func Success(ctx *fiber.Ctx, data interface{}) error { return resp.Success(ctx, data) }

func (r *Response) Success(ctx *fiber.Ctx, data interface{}) error {
	if data == nil {
		data = fiber.Map{}
	}

	return ctx.Status(fiber.StatusOK).JSON(Response{
		Code:    errcode.Success.Code(),
		Message: errcode.Success.Msg(),
		Data:    data,
	})
}

// Error return a error response
func Error(ctx *fiber.Ctx, err error) error { return resp.Error(ctx, err) }

func (r *Response) Error(ctx *fiber.Ctx, err error) error {
	if err == nil {
		return ctx.Status(fiber.StatusOK).JSON(Response{
			Code:    errcode.Success.Code(),
			Message: errcode.Success.Msg(),
			Data:    fiber.Map{},
		})
	}

	if v, ok := err.(*errcode.Error); ok {
		response := Response{
			Code:    v.Code(),
			Message: v.Msg(),
			Data:    fiber.Map{},
			Details: []string{},
		}
		details := v.Details()
		if len(details) > 0 {
			response.Details = details
		}
		return ctx.Status(v.StatusCode()).JSON(response)
	}
	return nil
}

// RouteNotFound 未找到相关路由
func RouteNotFound(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(Response{
		Code:    fiber.StatusNotFound,
		Message: "Not Found",
	})
}

// healthCheckResponse 健康检查响应结构体
type healthCheckResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

// HealthCheck will return OK if the underlying BoltDB is healthy. At least healthy enough for demoing purposes.
func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(healthCheckResponse{
		Status:   "UP",
		Hostname: ctx.Hostname(),
	})
}
