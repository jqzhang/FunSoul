package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"funWords": "三分天注定，七分靠打拼，还有90分看脸。",
	})
}
