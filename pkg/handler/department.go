package hendler

import (
	"github.com/Abdullayev65/attendance/pkg/ioPut"
	"github.com/Abdullayev65/attendance/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) DepAdd(c *gin.Context) {
	var depInput ioPut.DepInput
	c.Bind(&depInput)
	department := models.Department{Name: depInput.Name}
	err := h.Service.DepAdd(&department)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &department)
}
func (h *Handler) DepAll(c *gin.Context) {
	deps := h.Service.DepAll()
	c.JSON(http.StatusOK, deps)
}
