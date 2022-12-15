package servitorservices

import (
	"context"
	"errors"
	"servhunt/infra/utils"
	"servhunt/servitorservices/dao"
)

var (
	logger = utils.GetRootLogger()
)

type ServitorServices interface {
	CreateService(ctx context.Context, service ServiceRequest) (*ServiceResponse, error)
	UpdateService(ctx context.Context, service UpdateServiceRequest) (*ServiceResponse, error)
	CreateLocationInfo(ctx context.Context, service LocationInfoRequest) (*LocationInfoResponse, error)
	UpdateLocationInfo(ctx context.Context, service UpdateLocationInfoRequest) (*LocationInfoResponse, error)
	CreateCategory(ctx context.Context, category CategoryRequest) (*CategoryResponse, error)
	UpdateCategory(ctx context.Context, category UpdateCategoryRequest) (*CategoryResponse, error)
	ServiceLocations(ctx context.Context, serviceId int) (*[]LocationResponse, error)
	GetAllLocations(ctx context.Context) (*[]LocationResponse, error)
	ServiceCategories(ctx context.Context, serviceId int) (*[]CategoriesResponse, error)
	GetAllCategories(ctx context.Context) (*[]CategoriesResponse, error)
	GetAllServices(ctx context.Context) (*[]ServicesResponse, error)
	ServitorServices(ctx context.Context, userId int) (*[]ServicesResponse, error)
	GetServiceByID(ctx context.Context, id int) (*ServicesResponse, error)
}

type ServitorSvcImpl struct {
	dao.ServiceRepo
}

func NewServitorSvc(svc dao.ServiceRepo) ServitorServices {
	return &ServitorSvcImpl{ServiceRepo: svc}
}

func (s ServitorSvcImpl) CreateService(ctx context.Context, service ServiceRequest) (*ServiceResponse, error) {
	var finalLoc []dao.Location
	if service.Locations != nil {
		for _, loc := range service.Locations {
			locReq := dao.Location{
				LocationImage: loc.LocationImage,
				LocationName:  loc.LocationName,
				Latitude:      loc.Latitude,
				Longitude:     loc.Longitude,
				Address:       loc.Address,
			}
			finalLoc = append(finalLoc, locReq)
		}
	}
	var finalCat []dao.Category
	if service.Category != nil {
		for _, cat := range service.Category {
			catReq := dao.Category{
				CategoryImage: cat.CategoryImage,
				CategoryName:  cat.CategoryName,
			}
			finalCat = append(finalCat, catReq)
		}
	}

	svcReq := dao.Service{
		UserID:          service.UserID,
		ServiceImage:    service.ServiceImage,
		ServiceName:     service.ServiceName,
		ServiceDuration: service.ServiceDuration,
		ServiceCost:     service.ServiceCost,
		LocationInfo:    finalLoc,
		Category:        finalCat,
	}

	createService, err := s.ServiceRepo.CreateService(ctx, svcReq)
	if err != nil {
		return nil, err
	}

	res := ServiceResponse{ServiceId: createService.ID}

	return &res, nil

}

func (s ServitorSvcImpl) UpdateService(ctx context.Context, service UpdateServiceRequest) (*ServiceResponse, error) {
	svcReq := dao.Service{
		ServiceImage:    service.ServiceImage,
		ServiceName:     service.ServiceName,
		ServiceDuration: service.ServiceDuration,
		ServiceCost:     service.ServiceCost,
		ID:              service.ID,
	}
	updatedService, err := s.ServiceRepo.UpdateService(ctx, svcReq)
	if err != nil {
		return nil, err
	}
	res := ServiceResponse{ServiceId: updatedService.ID}

	return &res, nil
}

func (s ServitorSvcImpl) CreateLocationInfo(ctx context.Context, loc LocationInfoRequest) (*LocationInfoResponse, error) {
	locReq := dao.Location{
		LocationImage: loc.LocationImage,
		LocationName:  loc.LocationName,
		Latitude:      loc.Latitude,
		Longitude:     loc.Longitude,
		Address:       loc.Address,
		ServiceID:     loc.ServiceID,
	}
	locRes, err := s.ServiceRepo.CreateLocationInfo(ctx, locReq)
	if err != nil {
		return nil, err
	}
	res := LocationInfoResponse{LocationID: locRes.ID}
	return &res, nil
}

func (s ServitorSvcImpl) UpdateLocationInfo(ctx context.Context, loc UpdateLocationInfoRequest) (*LocationInfoResponse, error) {
	locReq := dao.Location{
		LocationImage: loc.LocationImage,
		LocationName:  loc.LocationName,
		Latitude:      loc.Latitude,
		Longitude:     loc.Longitude,
		Address:       loc.Address,
		ServiceID:     loc.ServiceID,
	}
	locRes, err := s.ServiceRepo.UpdateLocationInfo(ctx, locReq)
	if err != nil {
		return nil, err
	}
	res := LocationInfoResponse{LocationID: locRes.ID}
	return &res, nil
}

func (s ServitorSvcImpl) CreateCategory(ctx context.Context, category CategoryRequest) (*CategoryResponse, error) {
	catReq := dao.Category{
		CategoryImage: category.CategoryImage,
		CategoryName:  category.CategoryName,
		ServiceID:     category.ServiceID,
	}
	cat, err := s.ServiceRepo.CreateCategory(ctx, catReq)
	if err != nil {
		return nil, err
	}
	res := CategoryResponse{CategoryId: cat.ID}
	return &res, nil
}

func (s ServitorSvcImpl) UpdateCategory(ctx context.Context, category UpdateCategoryRequest) (*CategoryResponse, error) {
	catReq := dao.Category{
		CategoryImage: category.CategoryImage,
		CategoryName:  category.CategoryName,
		ServiceID:     category.ServiceID,
	}
	cat, err := s.ServiceRepo.UpdateCategory(ctx, catReq)
	if err != nil {
		return nil, err
	}
	res := CategoryResponse{CategoryId: cat.ID}
	return &res, nil
}

func (s ServitorSvcImpl) ServiceLocations(ctx context.Context, serviceId int) (*[]LocationResponse, error) {
	var locs []LocationResponse
	locations, err := s.ServiceRepo.ServiceLocations(ctx, serviceId)
	if err != nil {
		return nil, err
	}
	if len(*locations) < 1 {
		return nil, errors.New("no locations found")
	}
	for _, locss := range *locations {
		loc := LocationResponse{
			ID:            locss.ID,
			ServiceID:     locss.ServiceID,
			LocationImage: locss.LocationImage,
			LocationName:  locss.LocationName,
			Latitude:      locss.Latitude,
			Longitude:     locss.Longitude,
			Address:       locss.Address,
		}
		locs = append(locs, loc)
	}
	return &locs, nil
}

func (s ServitorSvcImpl) GetAllLocations(ctx context.Context) (*[]LocationResponse, error) {
	var locs []LocationResponse
	locations, err := s.ServiceRepo.GetAllLocations(ctx)
	if err != nil {
		return nil, err
	}
	if len(*locations) < 1 {
		return nil, errors.New("no locations found")
	}
	for _, locss := range *locations {
		loc := LocationResponse{
			ID:            locss.ID,
			ServiceID:     locss.ServiceID,
			LocationImage: locss.LocationImage,
			LocationName:  locss.LocationName,
			Latitude:      locss.Latitude,
			Longitude:     locss.Longitude,
			Address:       locss.Address,
		}
		locs = append(locs, loc)
	}
	return &locs, nil
}

func (s ServitorSvcImpl) ServiceCategories(ctx context.Context, serviceId int) (*[]CategoriesResponse, error) {
	var cats []CategoriesResponse
	categories, err := s.ServiceRepo.ServiceCategories(ctx, serviceId)
	if err != nil {
		return nil, err
	}
	if len(*categories) < 1 {
		return nil, errors.New("no categories found")
	}
	for _, cat := range *categories {
		catt := CategoriesResponse{
			ID:            cat.ID,
			ServiceID:     cat.ServiceID,
			CategoryImage: cat.CategoryImage,
			CategoryName:  cat.CategoryName,
		}
		cats = append(cats, catt)
	}
	return &cats, nil
}

func (s ServitorSvcImpl) GetAllCategories(ctx context.Context) (*[]CategoriesResponse, error) {
	var cats []CategoriesResponse
	categories, err := s.ServiceRepo.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}
	if len(*categories) < 1 {
		return nil, errors.New("no categories found")
	}
	for _, cat := range *categories {
		catt := CategoriesResponse{
			ID:            cat.ID,
			ServiceID:     cat.ServiceID,
			CategoryImage: cat.CategoryImage,
			CategoryName:  cat.CategoryName,
		}
		cats = append(cats, catt)
	}
	return &cats, nil
}

func (s ServitorSvcImpl) GetAllServices(ctx context.Context) (*[]ServicesResponse, error) {
	serviceList, err := s.ServiceRepo.GetAllServices(ctx)
	if err != nil {
		return nil, err
	}
	var resP []ServicesResponse
	var locs []Locations
	var cats []Category
	for _, svc := range *serviceList {
		if svc.LocationInfo != nil {
			for _, loc := range svc.LocationInfo {
				finalLoc := Locations{
					LocationImage: loc.LocationImage,
					LocationName:  loc.LocationName,
					Latitude:      loc.Latitude,
					Longitude:     loc.Longitude,
					Address:       loc.Address,
				}
				locs = append(locs, finalLoc)
			}
		}
		if svc.Category != nil {
			for _, cat := range svc.Category {
				finalCat := Category{
					CategoryImage: cat.CategoryImage,
					CategoryName:  cat.CategoryName,
				}
				cats = append(cats, finalCat)
			}
		}
		res := ServicesResponse{
			ID:              svc.ID,
			UserID:          svc.UserID,
			ServiceImage:    svc.ServiceImage,
			ServiceName:     svc.ServiceName,
			ServiceDuration: svc.ServiceDuration,
			ServiceCost:     svc.ServiceCost,
			Locations:       locs,
			Category:        cats,
		}
		resP = append(resP, res)
	}
	return &resP, nil
}

func (s ServitorSvcImpl) ServitorServices(ctx context.Context, userId int) (*[]ServicesResponse, error) {
	serviceList, err := s.ServiceRepo.ServitorServices(ctx, userId)
	if err != nil {
		return nil, err
	}
	var resP []ServicesResponse
	var locs []Locations
	var cats []Category
	for _, svc := range *serviceList {
		if svc.LocationInfo != nil {
			for _, loc := range svc.LocationInfo {
				finalLoc := Locations{
					LocationImage: loc.LocationImage,
					LocationName:  loc.LocationName,
					Latitude:      loc.Latitude,
					Longitude:     loc.Longitude,
					Address:       loc.Address,
				}
				locs = append(locs, finalLoc)
			}
		}
		if svc.Category != nil {
			for _, cat := range svc.Category {
				finalCat := Category{
					CategoryImage: cat.CategoryImage,
					CategoryName:  cat.CategoryName,
				}
				cats = append(cats, finalCat)
			}
		}
		res := ServicesResponse{
			ID:              svc.ID,
			UserID:          svc.UserID,
			ServiceImage:    svc.ServiceImage,
			ServiceName:     svc.ServiceName,
			ServiceDuration: svc.ServiceDuration,
			ServiceCost:     svc.ServiceCost,
			Locations:       locs,
			Category:        cats,
		}
		resP = append(resP, res)
	}
	return &resP, nil
}

func (s ServitorSvcImpl) GetServiceByID(ctx context.Context, id int) (*ServicesResponse, error) {
	svc, err := s.ServiceRepo.GetServiceByID(ctx, id)
	if err != nil {
		return nil, err
	}
	var locs []Locations
	var cats []Category
	if svc.LocationInfo != nil {
		for _, loc := range svc.LocationInfo {
			finalLoc := Locations{
				LocationImage: loc.LocationImage,
				LocationName:  loc.LocationName,
				Latitude:      loc.Latitude,
				Longitude:     loc.Longitude,
				Address:       loc.Address,
			}
			locs = append(locs, finalLoc)
		}
	}
	if svc.Category != nil {
		for _, cat := range svc.Category {
			finalCat := Category{
				CategoryImage: cat.CategoryImage,
				CategoryName:  cat.CategoryName,
			}
			cats = append(cats, finalCat)
		}
	}
	res := ServicesResponse{
		ID:              svc.ID,
		UserID:          svc.UserID,
		ServiceImage:    svc.ServiceImage,
		ServiceName:     svc.ServiceName,
		ServiceDuration: svc.ServiceDuration,
		ServiceCost:     svc.ServiceCost,
		Locations:       locs,
		Category:        cats,
	}
	return &res, nil
}
