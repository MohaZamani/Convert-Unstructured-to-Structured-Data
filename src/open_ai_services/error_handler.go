package openaiservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

func OpenAIErrorHandler(err openai.APIError, c *gin.Context) {
	switch err.HTTPStatusCode {
		case 401:
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "access token issue!",
				"open_ai_message": err.Message,
				"open_ai_status_code": err.HTTPStatusCode,
			})
		case 429:
			c.JSON(http.StatusForbidden, gin.H{
				"message": "rate limit issue, please try later!",
				"open_ai_message": err.Message,
				"open_ai_status_code": err.HTTPStatusCode,
			})
		case 500:
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "Third party code issue",
				"open_ai_message": err.Message,
				"open_ai_status_code": err.HTTPStatusCode,
			})
		default:
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": "unexpected issue!",
				"open_ai_message": err.Message,
				"open_ai_status_code": err.HTTPStatusCode,
			})
	}
}

