package routes

import (
	"net/http"
	"sqlc-crud-go/dbsqlc"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, queries *dbsqlc.Queries) {
	router.GET("/authors", func(c *gin.Context) {
		authors, err := queries.ListAuthors(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch authors"})
			return
		}
		c.JSON(http.StatusOK, authors)
	})

	router.GET("/authors/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		author, err := queries.GetAuthor(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch author"})
			return
		}
		c.JSON(http.StatusOK, author)
	})

	router.POST("/authors", func(c *gin.Context) {
		var input dbsqlc.CreateAuthorParams
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		author, err := queries.CreateAuthor(c, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create author"})
			return
		}

		c.JSON(http.StatusOK, author)
	})

	router.DELETE("/authors/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		err = queries.DeleteAuthor(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete author"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})

	})

}
