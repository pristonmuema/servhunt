package dao

import (
	"context"
	"gorm.io/gorm/clause"
	"servhunt/infra/dao"
	"time"
)

type ServiceRepo interface {
	CreateService(ctx context.Context, service Service) (*Service, error)
	UpdateService(ctx context.Context, service Service) (*Service, error)
	GetAllServices(ctx context.Context) (*[]Service, error)
	ServitorServices(ctx context.Context, userId int) (*[]Service, error)
	GetServiceByID(ctx context.Context, id int) (*Service, error)
	CreateCategory(ctx context.Context, category Category) (*Category, error)
	UpdateCategory(ctx context.Context, category Category) (*Category, error)
	ServiceCategories(ctx context.Context, serviceId int) (*[]Category, error)
	GetAllCategories(ctx context.Context) (*[]Category, error)
	CreateLocationInfo(ctx context.Context, location Location) (*Location, error)
	UpdateLocationInfo(ctx context.Context, location Location) (*Location, error)
	ServiceLocations(ctx context.Context, serviceId int) (*[]Location, error)
	GetAllLocations(ctx context.Context) (*[]Location, error)
}

type ServiceRepoImpl struct {
	repo *dao.Repository
}

func NewServiceRepoImpl(repo *dao.Repository) ServiceRepo {
	return &ServiceRepoImpl{repo: repo}
}

func (s *ServiceRepoImpl) CreateService(ctx context.Context, service Service) (*Service, error) {
	err := s.repo.DB.WithContext(ctx).Model(&Service{}).Create(&service).Error
	s.repo.DB.Save(&service)
	if err != nil {
		return nil, err
	}
	return &service, nil
}

func (s *ServiceRepoImpl) UpdateService(ctx context.Context, service Service) (*Service, error) {
	err := s.repo.DB.WithContext(ctx).Model(&Service{}).Where("id = ?", service.ID).Updates(Service{
		ServiceImage:    service.ServiceImage,
		ServiceName:     service.ServiceName,
		ServiceDuration: service.ServiceDuration,
		ServiceCost:     service.ServiceCost,
		LastUpdatedOn:   time.Now(),
	}).Error

	if err != nil {
		return nil, err
	}
	return &service, nil
}

func (s *ServiceRepoImpl) CreateCategory(ctx context.Context, category Category) (*Category, error) {
	err := s.repo.DB.WithContext(ctx).Model(&Category{}).Create(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *ServiceRepoImpl) UpdateCategory(ctx context.Context, category Category) (*Category, error) {
	err := s.repo.DB.WithContext(ctx).Model(&Service{}).Where("id = ?", category.ID).Updates(Category{
		CategoryImage: category.CategoryImage,
		CategoryName:  category.CategoryName,
		LastUpdatedOn: time.Time{},
	}).Error

	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *ServiceRepoImpl) CreateLocationInfo(ctx context.Context, location Location) (*Location, error) {
	err := s.repo.DB.WithContext(ctx).Model(&Location{}).Create(&location).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func (s *ServiceRepoImpl) UpdateLocationInfo(ctx context.Context, location Location) (*Location, error) {
	err := s.repo.DB.WithContext(ctx).Model(&Service{}).Where("id = ?", location.ID).Updates(Location{
		LocationImage: location.LocationImage,
		LocationName:  location.LocationName,
		Latitude:      location.Latitude,
		Longitude:     location.Longitude,
		Address:       location.Address,
		LastUpdatedOn: time.Now(),
	}).Error

	if err != nil {
		return nil, err
	}
	return &location, nil
}

func (s *ServiceRepoImpl) GetAllServices(ctx context.Context) (*[]Service, error) {
	var services []Service
	err := s.repo.DB.WithContext(ctx).Model(&Service{}).Preload(clause.Associations).Find(&services).Error
	if err != nil {
		return nil, err
	}
	return &services, nil
}

func (s *ServiceRepoImpl) ServitorServices(ctx context.Context, userId int) (*[]Service, error) {
	var services []Service
	err := s.repo.DB.WithContext(ctx).Model(&Service{}).Where("user_id = ?", userId).Preload(clause.Associations).Find(&services).Error
	if err != nil {
		return nil, err
	}
	return &services, nil
}

func (s *ServiceRepoImpl) GetServiceByID(ctx context.Context, id int) (*Service, error) {
	var services Service
	err := s.repo.DB.WithContext(ctx).Model(&Service{}).Where("id = ?", id).Preload(clause.Associations).Take(&services).Error
	if err != nil {
		return nil, err
	}
	return &services, nil
}

func (s *ServiceRepoImpl) ServiceCategories(ctx context.Context, serviceId int) (*[]Category, error) {
	var cats []Category
	err := s.repo.DB.WithContext(ctx).Model(&Category{}).Where("service_id = ?", serviceId).Preload(clause.Associations).Take(&cats).Error
	if err != nil {
		return nil, err
	}
	return &cats, nil
}

func (s *ServiceRepoImpl) GetAllCategories(ctx context.Context) (*[]Category, error) {
	var cats []Category
	err := s.repo.DB.WithContext(ctx).Model(&Category{}).Preload(clause.Associations).Find(&cats).Error
	if err != nil {
		return nil, err
	}
	return &cats, nil
}

func (s *ServiceRepoImpl) ServiceLocations(ctx context.Context, serviceId int) (*[]Location, error) {
	var locs []Location
	err := s.repo.DB.WithContext(ctx).Model(&Location{}).Where("service_id = ?", serviceId).Preload(clause.Associations).Take(&locs).Error
	if err != nil {
		return nil, err
	}
	return &locs, nil
}

func (s *ServiceRepoImpl) GetAllLocations(ctx context.Context) (*[]Location, error) {
	var locs []Location
	err := s.repo.DB.WithContext(ctx).Model(&Location{}).Preload(clause.Associations).Find(&locs).Error
	if err != nil {
		return nil, err
	}
	return &locs, nil
}
