package middleware

import (
    "crypto/subtle"
    "pplx2api/config"
    "strings"

    "github.com/gin-gonic/gin"
)

// AuthMiddleware validates Authorization header using constant-time compare
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if config.ConfigInstance.APIKey == "" {
            c.JSON(500, gin.H{
                "error": "API key not configured",
            })
            c.Abort()
            return
        }
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(401, gin.H{
                "error": "Missing Authorization header",
            })
            c.Abort()
            return
        }
        token = strings.TrimPrefix(token, "Bearer ")
        if subtle.ConstantTimeCompare([]byte(token), []byte(config.ConfigInstance.APIKey)) != 1 {
            c.JSON(401, gin.H{
                "error": "Invalid API key",
            })
            c.Abort()
            return
        }
        c.Next()
    }
}
