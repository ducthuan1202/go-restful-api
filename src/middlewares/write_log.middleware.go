package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func WriteLogApiRequest(c *gin.Context) {
	fmt.Printf("\n------ API MIDDLEWARE -----\n")
}

func WriteLogWebhookRequest(c *gin.Context) {
	fmt.Printf("\n------ WEBHOOK MIDDLEWARE -----\n")
}

func WriteLogRequestDetail(param gin.LogFormatterParams) string {
	// your custom format
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
