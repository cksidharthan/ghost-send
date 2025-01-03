package http

import (
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

		// If no password provided, render the password input form
		if password == "" {
			c.JSON(http.StatusOK, gin.H{"message": "Please enter a password"})
			return
		}

		result, err := secretsHandler.SecretsSvc.GetSecret(c.Request.Context(), db.GetSecretByIDParams{
			SecretID: uuid.MustParse(secretID),
			Password: password,
		})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}

		// Otherwise return full page
		c.JSON(http.StatusOK, SecretResponse{SecretText: result.SecretText})
	}
}
