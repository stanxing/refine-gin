package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stanxing/refine-gin/pkg/repository"
	"github.com/stanxing/refine-gin/pkg/resource"
)

// GenerateDeleteHandler generates a handler for DELETE operations
func GenerateDeleteHandler(res resource.Resource, repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get ID from URL parameters
		id := c.Param("id")

		// Call repository
		err := repo.Delete(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return result
		c.Status(http.StatusNoContent)
	}
}

// GenerateDeleteHandlerWithParam generates a handler for DELETE operations with custom ID parameter name
func GenerateDeleteHandlerWithParam(res resource.Resource, repo repository.Repository, idParamName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get ID from URL parameters using custom parameter name
		id := c.Param(idParamName)

		// Call repository
		err := repo.Delete(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return result
		c.Status(http.StatusNoContent)
	}
}
