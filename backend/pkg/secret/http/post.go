package http

import (
	"net/http"
	"time"

	db "github.com/cksidharthan/ghost-send/db/sqlc"
	"github.com/gin-gonic/gin"
)

type SecretRequest struct {
	SecretText string `json:"secret_text"`
	Password   string `json:"password"`
	Expiration string `json:"expiration"`
	Views      int    `json:"views"`
}

func postSecret(secretsHandler SecretHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request SecretRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate inputs
		if request.SecretText == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Secret text is required"})
			return
		}

		// Convert views to integer
		if request.Views < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid views value - must be a positive number"})
			return
		}

		// Convert expiration string to duration
		expirationDuration := getExpirationDuration(request.Expiration)

		result, err := secretsHandler.SecretsSvc.CreateSecret(c.Request.Context(), db.CreateSecretParams{
			SecretText:     request.SecretText,
			Password:       request.Password,
			ExpiresAt:      time.Now().Add(expirationDuration),
			RemainingTries: int32(request.Views),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Otherwise return full page
		c.JSON(http.StatusOK, gin.H{"message": "Secret created successfully", "secret_id": result})
	}
}

func getExpirationDuration(expiration string) time.Duration {
	switch expiration {
	case "5m":
		return 5 * time.Minute
	case "1h":
		return time.Hour
	case "1d":
		return 24 * time.Hour
	case "7d":
		return 7 * 24 * time.Hour
	default:
		return time.Hour
	}
}
