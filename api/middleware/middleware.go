package middleware

import (
	"api-gateway/api/token"
	"api-gateway/generated/user"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(auth user.AuthServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Missing access token"})
			return
		}

		ctx := context.Background()
		resp, err := auth.ValidateToken(ctx, &user.ValidateTokenReq{Token: tokenString})
		if err != nil {
			log.Println("Error validating token", err)
			c.AbortWithStatusJSON(403, gin.H{"error": "Invalid access token"})
			return
		}
		log.Println(resp.Valid, "Access token is valid")
		if !resp.Valid {
			c.AbortWithStatusJSON(403, gin.H{"error": "Invalid access token"})
			return
		}
		claims := token.Claims{
			Id:    resp.UserId,
			Email: resp.Email,
			Role:  resp.Role,
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func IsAuthorize(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing token claims",
			})
			c.Abort()
			return
		}

		claim, err := token.TokenClaimsParse(claims)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		log.Println(c.FullPath(), claim.GetRole(), c.Request.Method, "Authorization check")
		// Enforce checks whether the role has access to the resource
		ok, err := enforcer.Enforce(claim.GetRole(), c.FullPath(), c.Request.Method)
		fmt.Println(c.FullPath(), claim.GetRole(), c.Request.Method, ok, err)
		fmt.Println("Enforcer result: ", ok, " error: ", err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Authorization error",
			})
			c.Abort()
			return
		}

		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Sizda bu amalni bajarish uchun ruxsat yo'q.",
			})
			c.Abort()
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
