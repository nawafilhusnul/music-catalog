package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	membershipsmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
)

func (h *handler) SignUp(c *gin.Context) {
	var req membershipsmodel.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.svc.SignUp(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
