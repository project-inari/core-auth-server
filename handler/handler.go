// Package handler provides the handler functions for the domain
package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/project-inari/core-auth-server/service"
)

// Dependencies represents the dependencies required by the handler
type Dependencies struct {
	Service service.Port
}

// New creates a new handler
func New(e *echo.Echo, d Dependencies) *GRPCHandler {
	httpEcho := newHTTPHandler(d)
	httpEcho.initRoutes(e)

	return newGRPCHandler(d)
}
