package service

import (
	"errors"
	"github.com/Abdullayev65/attendance/pkg/models"
)

func (s *Service) AttendanceAdd(attendance *models.Attendance) error {
	if attendance.Type != 1 && attendance.Type != 0 {
		return errors.New("attendance type can be only 1 or 0")
	}
	err := s.DB.NewInsert().Model(attendance).Scan(s.ctx)
	return err
}
func (s *Service) AttendanceAll() *[]models.Attendance {
	all := new([]models.Attendance)
	s.DB.NewSelect().Model(all).Scan(s.ctx)
	return all
}
func (s *Service) AttendanceAllByUserID(userID int) (*[]models.Attendance, error) {
	all := new([]models.Attendance)
	err := s.DB.NewSelect().Model(all).
		Where("user_id = ?", userID).Scan(s.ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}
