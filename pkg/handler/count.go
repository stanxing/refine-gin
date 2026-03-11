package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stanxing/refine-gin/pkg/query"
	"github.com/stanxing/refine-gin/pkg/repository"
	"github.com/stanxing/refine-gin/pkg/resource"
)

// GenerateCountHandler generates a handler for COUNT operations
func GenerateCountHandler(res resource.Resource, repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		options := query.NewQueryOptions(c, res)
		options.DisablePagination = true

		count, err := repo.Count(c.Request.Context(), options)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"count": count,
		})
	}
}
