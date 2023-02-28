package hendler

import (
	"github.com/Abdullayev65/attendance/pkg/ioPut"
	"github.com/Abdullayev65/attendance/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) PositionAdd(c *gin.Context) {
	var positionInput ioPut.PositionInput
	c.Bind(&positionInput)
	department := models.Position{Name: positionInput.Name}
	err := h.Service.PositionAdd(&department)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &department)
}
func (h *Handler) PositionAll(c *gin.Context) {
	positions := h.Service.PositionAll()
	c.JSON(http.StatusOK, positions)
}
