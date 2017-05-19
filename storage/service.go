package storage

import (
	"twreporter.org/go-api/models"
)

// GetService ...
func (g *GormMembershipStorage) GetService(name string) (models.Service, error) {
	var s models.Service
	err := g.db.First(&s, "name = ?", name).Error
	return s, err
}

// CreateService this func will create a service record
func (g *GormMembershipStorage) CreateService(json models.ServiceJSON) (models.Service, error) {
	var err error

	service := models.Service{
		Name: json.Name,
	}

	err = g.db.Create(&service).Error

	return service, err
}

// UpdateService this func will update the record in the stroage
func (g *GormMembershipStorage) UpdateService(name string, json models.ServiceJSON) (models.Service, error) {
	var s models.Service
	var err error

	err = g.db.Where("name = ?", name).FirstOrCreate(&s).Error
	if err != nil {
		return s, err
	}

	s.Name = json.Name

	err = g.db.Save(&s).Error

	return s, err
}

// DeleteService this func will delete the record in the stroage
func (g *GormMembershipStorage) DeleteService(name string) error {
	err := g.db.Where("name = ?", name).Delete(&models.Service{}).Error
	return err
}