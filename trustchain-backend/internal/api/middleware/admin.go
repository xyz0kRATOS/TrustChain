package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/amanp/trustchain-backend/internal/models"
)

func AdminOnly(adminWallet string) gin.HandlerFunc {
	normalizedAdmin := strings.ToLower(strings.TrimSpace(adminWallet))

	return func(c *gin.Context) {
		headerWallet := strings.ToLower(strings.TrimSpace(c.GetHeader("X-Admin-Wallet")))
		if normalizedAdmin == "" || headerWallet == "" || !strings.EqualFold(headerWallet, normalizedAdmin) {
			msg := "forbidden"
			c.AbortWithStatusJSON(http.StatusForbidden, models.APIResponse[any]{Data: nil, Error: &msg})
			return
		}
		c.Next()
	}
}
