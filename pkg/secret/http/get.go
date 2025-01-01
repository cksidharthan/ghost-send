package http

import (
	"bytes"
	"io"
	"net/http"

	db "github.com/cksidharthan/share-secret/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getSecret(secretsHandler SecretHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		secretID := c.Param("id")
		password := c.Query("password")

		// If no password provided, render the password input form
		if password == "" {
			html, err := secretsHandler.Templates.Open("secret-view.html")
			if err != nil {
				c.HTML(http.StatusInternalServerError, "error.html", err.Error())
				return
			}
			defer html.Close()

			buf := new(bytes.Buffer)
			_, err = io.Copy(buf, html)
			if err != nil {
				c.HTML(http.StatusInternalServerError, "error.html", err.Error())
				return
			}

			c.HTML(http.StatusOK, buf.String(), gin.H{
				"ID": secretID,
			})
			return
		}

		result, err := secretsHandler.SecretsSvc.GetSecret(c.Request.Context(), db.GetSecretByIDParams{
			SecretID: uuid.MustParse(secretID),
			Password: password,
		})
		if err != nil {
			c.HTML(http.StatusOK, "error.html", err.Error())
			return
		}

		// If HTMX request, return partial template
		if c.GetHeader("HX-Request") == "true" {
			c.HTML(http.StatusOK, "secret-result.html", result)
			return
		}

		// Otherwise return full page
		c.HTML(http.StatusOK, "index.html", result)
	}
}
