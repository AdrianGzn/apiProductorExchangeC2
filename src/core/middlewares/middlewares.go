package middlewares

import "github.com/gin-gonic/gin"
import "log"

func NewCorsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println("CORS middleware executed") // Agrega este log
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        if c.Request.Method == "OPTIONS" {
            log.Println("Handling OPTIONS request")
            c.AbortWithStatus(200)
            return
        }

        c.Next()
    }
}