package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, code int) {
	currentTime := time.Now().UTC().Format(time.RFC3339)
	currentPath := c.Request.URL.Path

	switch code {
	case 400:
		c.JSON(http.StatusBadRequest, HttpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusBadRequest,
			Message: "Bad request",
		})
	case 401:
		c.JSON(http.StatusUnauthorized, HttpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized access",
		})
	case 403:
		c.JSON(http.StatusForbidden, HttpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusForbidden,
			Message: "Forbidden access",
		})
	case 404:
		c.JSON(http.StatusNotFound, HttpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusNotFound,
			Message: "Resource not found",
		})
	case 405:
		c.JSON(http.StatusMethodNotAllowed, HttpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusMethodNotAllowed,
			Message: "Request method is not allowed",
		})
	case 429:
		c.JSON(http.StatusTooManyRequests, HttpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusTooManyRequests,
			Message: "Too many requests in a short time",
		})
	case 500:
		c.JSON(http.StatusInternalServerError, HttpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusInternalServerError,
			Message: "Server internal error occurred",
		})
	case 502:
		c.JSON(http.StatusBadGateway, HttpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusBadGateway,
			Message: "Upstream server is not available",
		})
	case 503:
		c.JSON(http.StatusServiceUnavailable, HttpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusServiceUnavailable,
			Message: "Service is not available",
		})
	case 504:
		c.JSON(http.StatusGatewayTimeout, HttpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusGatewayTimeout,
			Message: "Server gateway timeout",
		})
	default:
		c.JSON(http.StatusSeeOther, HttpResponse{
			Error:   true,
			Path:    currentPath,
			Time:    currentTime,
			Status:  http.StatusSeeOther,
			Message: "Unknown error occurred",
		})
	}

	c.Abort()
}
