package middleware

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

func DomainValidatorMiddleware(domain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		if colonIndex := strings.LastIndex(host, ":"); colonIndex != -1 {
			host, _, _ = net.SplitHostPort(c.Request.Host)
		}

		// Allow direct IP access even when a domain is configured.
		// This keeps domain-based access working while removing the hard reverse-proxy-only restriction.
		if host != domain && net.ParseIP(host) == nil {
			c.AbortWithStatus(403)
			return
		}

		c.Next()
	}
}
