package service

import "github.com/Abdullayev65/attendance/pkg/models"

func (s *Service) UserByUsername(username string) (*models.User, error) {
	user := new(models.User)
	err := s.DB.NewSelect().Model(user).
		Where("username = ?", username).Scan(s.ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) UserAdd(user *models.User) error {
	_, err := s.DB.NewInsert().Model(user).Exec(s.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UserAllFullData() *[]models.User {
	all := make([]models.User, 0)
	s.DB.NewSelect().Model(&all).
		Relation("Department").
		Relation("Position").
		Relation("Attendances").
		Scan(s.ctx)
	return &all
}
