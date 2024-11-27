package tracks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) SearchTracks(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query is required"})
		return
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		pageSize = 10
	}

	pageIndex, err := strconv.Atoi(c.Query("page_index"))
	if err != nil {
		pageIndex = 1
	}

	tracks, err := h.svc.Search(c.Request.Context(), query, pageSize, pageIndex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tracks)
}
