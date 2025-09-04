package middleware

import (
    "os"
    "strings"

    "github.com/gin-gonic/gin"
)

// CORSMiddleware handles CORS headers with optional origin whitelist via ALLOWED_ORIGINS
func CORSMiddleware() gin.HandlerFunc {
    allowed := strings.TrimSpace(os.Getenv("ALLOWED_ORIGINS"))
    var whitelist map[string]struct{}
    wildcard := true
    if allowed != "" {
        entries := strings.Split(allowed, ",")
        whitelist = make(map[string]struct{}, len(entries))
        wildcard = false
        for _, e := range entries {
            v := strings.TrimSpace(e)
            if v == "*" {
                wildcard = true
            }
            if v != "" {
                whitelist[v] = struct{}{}
            }
        }
    }
    return func(c *gin.Context) {
        origin := c.GetHeader("Origin")
        if wildcard {
            c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        } else if origin != "" {
            if _, ok := whitelist[origin]; ok {
                c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
                c.Writer.Header().Set("Vary", "Origin")
            }
        }
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}
