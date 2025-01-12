package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/project-inari/core-auth-server/dto"
	"github.com/project-inari/core-auth-server/pkg/request"
	"github.com/project-inari/core-auth-server/pkg/response"
)

type httpHandler struct {
	d Dependencies
}

func newHTTPHandler(d Dependencies) *httpHandler {
	return &httpHandler{
		d: d,
	}
}

// SignUp handles the sign up request
func (h *httpHandler) SignUp(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(dto.SignUpReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [SignUp] bad request: %v", err), "")
	}

	header := dto.SignUpReqHeader{
		AcceptLocale: c.Request().Header.Get("Accept-Locale"),
	}

	res, err := h.d.Service.SignUp(ctx, *req, header)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [SignUp] unable to sign up: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
