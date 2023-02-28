package service

import (
	"errors"
	"github.com/Abdullayev65/attendance/pkg/models"
)

func (s *Service) PositionAdd(position *models.Position) error {
	if position.Name == "" {
		return errors.New("name of position can not be null or blank")
	}
	_, err := s.DB.NewInsert().Model(position).Exec(s.ctx)
	return err
}
func (s *Service) PositionAll() *[]models.Position {
	positions := new([]models.Position)
	s.DB.NewSelect().Model(positions).Scan(s.ctx)
	return positions
}
