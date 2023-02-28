package hendler

import (
	"github.com/Abdullayev65/attendance/pkg/ioPut"
	"github.com/Abdullayev65/attendance/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func (h *Handler) AttendanceAdd(c *gin.Context) {
	input := new(ioPut.AttendanceInput)
	e := c.Bind(input)
	_ = e
	if input.Type == nil {
		c.String(400, "type can not be null")
		return
	}
	userID := h.getUserID(c)
	attendance := &models.Attendance{Type: *input.Type, Time: time.Now(), UserID: userID}
	err := h.Service.AttendanceAdd(attendance)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}
	c.JSON(200, attendance)
}
func (h *Handler) AttendanceAll(c *gin.Context) {
	all := h.Service.AttendanceAll()
	c.JSON(200, all)
}
func (h *Handler) AttendanceList(c *gin.Context) {
	usersFullData := h.Service.UserAllFullData()
	infos := make([]ioPut.AttendanceInfo, 0)
	for _, user := range *usersFullData {
		info := ioPut.AttendanceInfo{}
		if user.Attendances != nil {
			info.Attendances = user.Attendances
		}
		if user.Department != nil {
			info.Department = user.Department.Name
		}
		if user.Position != nil {
			info.Position = user.Position.Name
		}
		info.FullName = strings.Join([]string{user.FirstName, user.MiddleName, user.LastName}, " ")
		info.Username = user.Username
		info.Department = user.Department.Name
		infos = append(infos, info)
	}
	c.JSON(200, infos)
}
func (h *Handler) AttendanceByUserID(c *gin.Context) {
	userId := h.getUserID(c)
	all, err := h.Service.AttendanceAllByUserID(userId)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, all)
}
