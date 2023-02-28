package ioPut

import "github.com/Abdullayev65/attendance/pkg/models"

type AttendanceInfo struct {
	FullName    string
	Username    string
	Department  string
	Position    string
	Attendances []models.Attendance
}
