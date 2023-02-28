package ioPut

import "github.com/Abdullayev65/attendance/pkg/models"

type Sign struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserAdd struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	FirstName    string `bun:",nullzero" json:"firstName"`
	LastName     string `bun:",nullzero" json:"lastName"`
	MiddleName   string `bun:",nullzero" json:"middleName"`
	DepartmentID int    `json:"departmentID"`
	PositionID   int    `json:"positionID"`
}
type DepInput struct {
	Name string `json:"name"`
}
type PositionInput struct {
	Name string `json:"name"`
}
type AttendanceInput struct {
	Type *int `json:"type"`
}

func (ua *UserAdd) MapToUser() *models.User {
	return &models.User{
		Username:     ua.Username,
		Password:     ua.Password,
		FirstName:    ua.FirstName,
		LastName:     ua.LastName,
		MiddleName:   ua.MiddleName,
		DepartmentID: ua.DepartmentID,
		PositionID:   ua.PositionID,
	}
}
func (ua *UserAdd) Valid() (valid bool, fieldName string) {
	if ua.Username == "" {
		fieldName = "username"
		return
	}
	if ua.Password == "" {
		fieldName = "password"
		return
	}
	if ua.PositionID == 0 {
		fieldName = "positionID"
		return
	}
	if ua.DepartmentID == 0 {
		fieldName = "departmentID"
		return
	}
	return true, ""
}
