package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/pingTime", func(ctx *gin.Context) {
		// JSON serializer is available on gin context
		ctx.JSON(http.StatusOK, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})
	r.Run(":8000") // listen and serve on

}
