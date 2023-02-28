package service

import (
	"errors"
	"github.com/Abdullayev65/attendance/pkg/models"
)

func (s *Service) DepAdd(dep *models.Department) error {
	if dep.Name == "" {
		return errors.New("name of department can not be null or blank")
	}
	_, err := s.DB.NewInsert().Model(dep).Exec(s.ctx)
	return err
}
func (s *Service) DepAll() *[]models.Department {
	deps := new([]models.Department)
	s.DB.NewSelect().Model(deps).Scan(s.ctx)
	return deps
}
