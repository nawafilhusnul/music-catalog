package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	membershipsmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
)

func (h *handler) SignIn(c *gin.Context) {
	var req membershipsmodel.SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.svc.SignIn(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
