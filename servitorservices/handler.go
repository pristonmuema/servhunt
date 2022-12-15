package servitorservices

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"servhunt/infra/utils"
	"strconv"
)

type ServitorServicesHandler interface {
	CreateService(ctx *gin.Context)
	UpdateService(ctx *gin.Context)
	CreateLocationInfo(ctx *gin.Context)
	UpdateLocationInfo(ctx *gin.Context)
	CreateCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	GetAllServices(ctx *gin.Context)
	GetAllLocations(ctx *gin.Context)
	GetAllCategories(ctx *gin.Context)
	GetServiceLocations(ctx *gin.Context)
	GetServiceCategories(ctx *gin.Context)
	GetServiceByID(ctx *gin.Context)
	ServitorsService(ctx *gin.Context)
}

type ServitorServicesHandlerImpl struct {
	ServitorServices
}

func NewServitorServicesHandlerImpl(svc ServitorServices) ServitorServicesHandler {
	return &ServitorServicesHandlerImpl{ServitorServices: svc}
}

func (s *ServitorServicesHandlerImpl) CreateService(ctx *gin.Context) {
	req := ServiceRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}

	service, err := s.ServitorServices.CreateService(ctx, req)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if service.ServiceId > 0 {
		utils.APIResponse(ctx, "Service created successfully", http.StatusOK, true, service)
		return
	}
	utils.APIResponse(ctx, "Failed to create service", http.StatusBadRequest, false, nil)
}

func (s *ServitorServicesHandlerImpl) UpdateService(ctx *gin.Context) {

	fetchReq := FetchByIdRequest{}
	id, err := strconv.Atoi(ctx.Param("service_id"))
	if err != nil {
		utils.APIResponse(ctx, "Failed to convert string to int", http.StatusBadRequest,
			false, err.Error())
		return
	}
	fetchReq.ServiceId = id
	if err := ctx.ShouldBindJSON(&fetchReq); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}

	req := UpdateServiceRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}
	req.ID = fetchReq.ServiceId
	service, err := s.ServitorServices.UpdateService(ctx, req)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if service.ServiceId > 0 {
		utils.APIResponse(ctx, "Service updated successfully", http.StatusOK, true, service)
		return
	}
	utils.APIResponse(ctx, "Failed to update service", http.StatusBadRequest, false, nil)
}

func (s *ServitorServicesHandlerImpl) CreateLocationInfo(ctx *gin.Context) {
	req := LocationInfoRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}

	location, err := s.ServitorServices.CreateLocationInfo(ctx, req)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if location.LocationID > 0 {
		utils.APIResponse(ctx, "Location created successfully", http.StatusOK, true, location)
		return
	}
	utils.APIResponse(ctx, "Failed to create location", http.StatusBadRequest, false, nil)
}

func (s *ServitorServicesHandlerImpl) UpdateLocationInfo(ctx *gin.Context) {

	fetchReq := FetchLocByIdRequest{}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.APIResponse(ctx, "Failed to convert string to int", http.StatusBadRequest,
			false, err.Error())
		return
	}
	fetchReq.LocationID = id
	if err := ctx.ShouldBindJSON(&fetchReq); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}

	req := UpdateLocationInfoRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}
	req.ID = fetchReq.LocationID
	location, err := s.ServitorServices.UpdateLocationInfo(ctx, req)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if location.LocationID > 0 {
		utils.APIResponse(ctx, "Location updated successfully", http.StatusOK, true, location)
		return
	}
	utils.APIResponse(ctx, "Failed to update location", http.StatusBadRequest, false, nil)
}

func (s *ServitorServicesHandlerImpl) CreateCategory(ctx *gin.Context) {
	req := CategoryRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}

	category, err := s.ServitorServices.CreateCategory(ctx, req)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if category.CategoryId > 0 {
		utils.APIResponse(ctx, "Category created successfully", http.StatusOK, true, category)
		return
	}
	utils.APIResponse(ctx, "Failed to create category", http.StatusBadRequest, false, nil)
}

func (s *ServitorServicesHandlerImpl) UpdateCategory(ctx *gin.Context) {

	fetchReq := FetchCatByIdRequest{}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.APIResponse(ctx, "Failed to convert string to int", http.StatusBadRequest,
			false, err.Error())
		return
	}
	fetchReq.CategoryId = id
	if err := ctx.ShouldBindJSON(&fetchReq); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}

	req := UpdateCategoryRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}

	category, err := s.ServitorServices.UpdateCategory(ctx, req)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if category.CategoryId > 0 {
		utils.APIResponse(ctx, "Category updated successfully", http.StatusOK, true, category)
		return
	}
	utils.APIResponse(ctx, "Failed to update  category", http.StatusBadRequest, false, nil)
}

func (s *ServitorServicesHandlerImpl) GetAllServices(ctx *gin.Context) {
	users, err := s.ServitorServices.GetAllServices(ctx)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if len(*users) > 0 {
		utils.APIResponse(ctx, "Services successfully returned", http.StatusOK, true, users)
		return
	}
	utils.APIResponse(ctx, "Failed to return services", http.StatusBadRequest, false, nil)
}

func (s *ServitorServicesHandlerImpl) GetAllLocations(ctx *gin.Context) {
	locations, err := s.ServitorServices.GetAllLocations(ctx)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if len(*locations) > 0 {
		utils.APIResponse(ctx, "Locations successfully returned", http.StatusOK, true, locations)
		return
	}
	utils.APIResponse(ctx, "Failed to return locations", http.StatusBadRequest, false, nil)
}

func (s *ServitorServicesHandlerImpl) GetAllCategories(ctx *gin.Context) {
	categories, err := s.ServitorServices.GetAllCategories(ctx)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if len(*categories) > 0 {
		utils.APIResponse(ctx, "Categories successfully returned", http.StatusOK, true, categories)
		return
	}
	utils.APIResponse(ctx, "Failed to return categories", http.StatusBadRequest, false, nil)
}

func (s *ServitorServicesHandlerImpl) GetServiceLocations(ctx *gin.Context) {
	fetchReq := FetchByIdRequest{}
	id, err := strconv.Atoi(ctx.Param("service_id"))
	if err != nil {
		utils.APIResponse(ctx, "Failed to convert string to int", http.StatusBadRequest,
			false, err.Error())
		return
	}
	fetchReq.ServiceId = id
	if err := ctx.ShouldBindJSON(&fetchReq); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}

	categories, err := s.ServitorServices.ServiceLocations(ctx, fetchReq.ServiceId)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if len(*categories) > 0 {
		utils.APIResponse(ctx, "Service locations successfully returned", http.StatusOK, true, categories)
		return
	}
	utils.APIResponse(ctx, "Failed to return a service categories", http.StatusBadRequest, false, nil)
}

func (s *ServitorServicesHandlerImpl) GetServiceCategories(ctx *gin.Context) {

	fetchReq := FetchByIdRequest{}
	id, err := strconv.Atoi(ctx.Param("service_id"))
	if err != nil {
		utils.APIResponse(ctx, "Failed to convert string to int", http.StatusBadRequest,
			false, err.Error())
		return
	}
	fetchReq.ServiceId = id
	if err := ctx.ShouldBindJSON(&fetchReq); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}

	categories, err := s.ServitorServices.ServiceCategories(ctx, fetchReq.ServiceId)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if len(*categories) > 0 {
		utils.APIResponse(ctx, "Service categories successfully returned", http.StatusOK, true, categories)
		return
	}
	utils.APIResponse(ctx, "Failed to return service categories", http.StatusBadRequest, false, nil)
}

func (s *ServitorServicesHandlerImpl) GetServiceByID(ctx *gin.Context) {
	fetchReq := FetchByIdRequest{}
	id, err := strconv.Atoi(ctx.Param("service_id"))
	if err != nil {
		utils.APIResponse(ctx, "Failed to convert string to int", http.StatusBadRequest,
			false, err.Error())
		return
	}
	fetchReq.ServiceId = id
	if err := ctx.ShouldBindJSON(&fetchReq); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}

	category, err := s.ServitorServices.GetServiceByID(ctx, fetchReq.ServiceId)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if category.ID > 0 {
		utils.APIResponse(ctx, "Service successfully returned", http.StatusOK, true, category)
		return
	}
	utils.APIResponse(ctx, "Failed to return service", http.StatusBadRequest, false, nil)
}

func (s *ServitorServicesHandlerImpl) ServitorsService(ctx *gin.Context) {
	fetchReq := OrderIdRequest{}
	id, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		utils.APIResponse(ctx, "Failed to convert string to int", http.StatusBadRequest,
			false, err.Error())
		return
	}
	fetchReq.UserId = id
	if err := ctx.ShouldBindJSON(&fetchReq); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}
	svc, err := s.ServitorServices.ServitorServices(ctx, fetchReq.UserId)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if len(*svc) > 0 {
		utils.APIResponse(ctx, "Services successfully returned", http.StatusOK, true, svc)
		return
	}
	utils.APIResponse(ctx, "Failed to return services", http.StatusBadRequest, false, nil)
}
