package middleware

import (
	"api-gateway/api/token"
	"api-gateway/generated/user"
	"context"
	"log/slog"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(authServiceClient user.AuthServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Missing access token"})
			return
		}

		ctx := context.Background()
		resp, err := authServiceClient.ValidateToken(ctx, &user.ValidateTokenReq{Token: tokenString})
		if err != nil {
			c.AbortWithStatusJSON(403, gin.H{"error": "Invalid access token"})
			return
		}

		if !resp.Valid {
			c.AbortWithStatusJSON(403, gin.H{"error": "Invalid access token"})
			return
		}

		c.Set("claims", token.Claims{
			Id:    resp.UserId,
			Email: resp.Email,
			Role:  resp.Role,
		})
		c.Next()
	}
}

func IsAuthorize(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token claims"})
			return
		}

		claim, err := token.TokenClaimsParse(claims)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		// Enforce checks whether the role has access to the resource
		ok, err := enforcer.Enforce(claim.GetRole(), c.FullPath(), c.Request.Method)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Authorization error"})
			return
		}

		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			return
		}

		c.Next()
	}
}

func LogMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Request received",
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
		)

		c.Next()

		logger.Info("Response sent",
			slog.Int("status", c.Writer.Status()),
		)
	}
}
