package http

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"

	db "github.com/cksidharthan/share-secret/db/sqlc"
	"github.com/gin-gonic/gin"
)

func postSecret(secretsHandler SecretHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", err.Error())
			return
		}

		// Get form values
		secretText := c.PostForm("secret")
		password := c.PostForm("password")
		expiration := c.PostForm("expiration")
		viewsStr := c.PostForm("views")

		// Validate inputs
		if secretText == "" {
			c.HTML(http.StatusBadRequest, "error.html", "Secret text is required")
			return
		}

		// Convert views to integer
		maxViews, err := strconv.Atoi(viewsStr)
		if err != nil || maxViews < 1 {
			c.HTML(http.StatusBadRequest, "error.html", "Invalid views value - must be a positive number")
			return
		}

		// Convert expiration string to duration
		expirationDuration := getExpirationDuration(expiration)

		result, err := secretsHandler.SecretsSvc.CreateSecret(c.Request.Context(), db.CreateSecretParams{
			SecretText:     secretText,
			Password:       password,
			ExpiresAt:      time.Now().Add(expirationDuration),
			RemainingTries: int32(maxViews),
		})
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", err.Error())
			return
		}

		// If HTMX request, return partial template
		if c.GetHeader("HX-Request") == "true" {
			html, err := secretsHandler.Templates.ReadFile("templates/secret-result.html")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read secret-result.html file"})
				return
			}

			fullURL := fmt.Sprintf("%s%s", "https://secret.cksidharthan.site/secret/", result)

			html = bytes.Replace(html, []byte("{{.}}"), []byte(fullURL), 2)

			c.DataFromReader(http.StatusOK, int64(len(html)), "text/html", bytes.NewReader(html), nil)
			return
		}

		// Otherwise return full page
		c.HTML(http.StatusOK, "index.html", result)
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
