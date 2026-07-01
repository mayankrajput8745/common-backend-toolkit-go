package httpserver

import (
	"net/http"

	"github.com/mayankrajput8745/common-backend-toolkit-go/httpserver/contract"
)

type Response struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type Pagination struct {
	TotalCount int    `json:"total_count"`
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	NextCursor string `json:"next_cursor,omitempty"`
	PrevCursor string `json:"prev_cursor,omitempty"`
}

type PaginatedDataResponse struct {
	Success    bool       `json:"success"`
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func SuccessResponse(ctx contract.Context, data any) error {
	return ctx.WriteJSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func SuccessMsgResponse(ctx contract.Context, msg string) error {
	return ctx.WriteJSON(http.StatusAccepted, Response{
		Success: true,
		Msg:     msg,
	})
}

func CreatedResponse(ctx contract.Context, data any) error {
	return ctx.WriteJSON(http.StatusCreated, Response{
		Success: true,
		Data:    data,
	})
}

func InternalServerErrorResponse(ctx contract.Context, msg string) error {
	return ctx.WriteJSON(http.StatusInternalServerError, Response{
		Success: false,
		Msg:     msg,
	})
}

func BadRequestResponse(ctx contract.Context, msg string) error {
	return ctx.WriteJSON(http.StatusBadRequest, Response{
		Success: false,
		Msg:     msg,
	})
}

func NotFoundResponse(ctx contract.Context, msg string) error {
	return ctx.WriteJSON(http.StatusNotFound, Response{
		Success: false,
		Msg:     msg,
	})
}

func UnauthorizedResponse(ctx contract.Context, msg string) error {
	return ctx.WriteJSON(http.StatusUnauthorized, Response{
		Success: false,
		Msg:     msg,
	})
}

func ForbiddenResponse(ctx contract.Context, msg string) error {
	return ctx.WriteJSON(http.StatusForbidden, Response{
		Success: false,
		Msg:     msg,
	})
}

func TooManyRequestsResponse(ctx contract.Context, msg string) error {
	return ctx.WriteJSON(http.StatusTooManyRequests, Response{
		Success: false,
		Msg:     msg,
	})
}

func PaginatedResponse(ctx contract.Context, data any, pagination Pagination) error {
	return ctx.WriteJSON(http.StatusOK, PaginatedDataResponse{
		Success:    true,
		Data:       data,
		Pagination: pagination,
	})
}

func CustomResponse(ctx contract.Context, statusCode int, success bool, msg string, data any) error {
	return ctx.WriteJSON(statusCode, Response{
		Success: success,
		Msg:     msg,
		Data:    data,
	})
}

func TextResponse(ctx contract.Context, text string) error {
	ctx.SetHeader("Content-Type", "text/plain")
	return ctx.String(text)
}
