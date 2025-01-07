package http

import (
	"errors"
	"github.com/cksidharthan/share-secret/pkg/secret/svc"
	"net/http"

	db "github.com/cksidharthan/share-secret/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SecretResponse struct {
	SecretText string `json:"secret_text"`
}

func getSecret(secretsHandler SecretHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		secretID := c.Param("id")
		password := c.Query("password")

		parsedUUID, err := uuid.Parse(secretID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid secret ID"})
			return
		}

		// If no password provided, render the password input form
		if password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Please enter a password"})
			return
		}

		result, err := secretsHandler.SecretsSvc.GetSecret(c.Request.Context(), db.GetSecretByIDParams{
			SecretID: parsedUUID,
			Password: password,
		})
		if err != nil {
			if errors.Is(err, svc.ErrInvalidPassword) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// Otherwise return full page
		c.JSON(http.StatusOK, SecretResponse{SecretText: result.SecretText})
	}
}
