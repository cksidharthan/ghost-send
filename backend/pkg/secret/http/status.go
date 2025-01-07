package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getSecretStatus(secretsHandler SecretHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		secretID := c.Param("id")

		parsedUUID, err := uuid.Parse(secretID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		result, err := secretsHandler.SecretsSvc.CheckSecretExists(c.Request.Context(), parsedUUID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		if !result {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusOK)
	}
}
