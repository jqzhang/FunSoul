package middleware

import "github.com/gin-gonic/gin"

func Set() gin.HandlerFunc {

	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
	}
}
