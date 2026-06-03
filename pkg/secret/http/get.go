package http

import (
	"errors"
	"net/http"

	"github.com/cksidharthan/ghost-send/pkg/secret/svc"

	db "github.com/cksidharthan/ghost-send/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SecretResponse struct {
	SecretText string `json:"secret_text"`
}

type GetSecretRequest struct {
	Password string `json:"password"`
}

func getSecret(secretsHandler SecretHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		secretID := c.Param("id")

		parsedUUID, err := uuid.Parse(secretID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid secret ID"})
			return
		}

		var request GetSecretRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if request.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Please enter a password"})
			return
		}

		result, err := secretsHandler.SecretsSvc.GetSecret(c.Request.Context(), db.GetSecretByIDLockedParams{
			SecretID: parsedUUID,
			Password: request.Password,
		})
		if err != nil {
			if errors.Is(err, svc.ErrInvalidPassword) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
				return
			}
			if errors.Is(err, svc.ErrSecretNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Secret not found or expired"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, SecretResponse{SecretText: result.SecretText})
	}
}
