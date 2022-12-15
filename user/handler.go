package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"servhunt/infra/utils"
	"strconv"
)

type UsersHandler interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
	CreateUserAccount(ctx *gin.Context)
	UpdateUserAccount(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	GetUserByPhoneNo(ctx *gin.Context)
	GetUserByEmail(ctx *gin.Context)
}

type UsersHandlerImpl struct {
	UserService
}

func NewUsersHandlerImpl(svc UserService) UsersHandler {
	return &UsersHandlerImpl{UserService: svc}
}

func (user *UsersHandlerImpl) Login(ctx *gin.Context) {
	req := LoginRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}
	login, err := user.UserService.Login(ctx, req)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if login.AccessToken != "" {
		utils.APIResponse(ctx, "Access granted", http.StatusOK, true, login)
		return
	}
	utils.APIResponse(ctx, "Access denied", http.StatusBadRequest, false, nil)
}

func (user *UsersHandlerImpl) Logout(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (user *UsersHandlerImpl) ChangePassword(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (user *UsersHandlerImpl) CreateUserAccount(ctx *gin.Context) {
	req := CreateUserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}

	account, err := user.UserService.CreateUserAccount(ctx, req)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if account.UserId > 0 {
		utils.APIResponse(ctx, "User account created successfully", http.StatusOK, true, account)
		return
	}
	utils.APIResponse(ctx, "Failed to create user account", http.StatusBadRequest, false, nil)
}

func (user *UsersHandlerImpl) UpdateUserAccount(ctx *gin.Context) {

	fetchReq := FetchByIdRequest{}
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

	req := UpdateUserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}
	req.ID = fetchReq.UserId
	account, err := user.UserService.UpdateUserAccount(ctx, req)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if account.UserId > 0 {
		utils.APIResponse(ctx, "User account updated successfully", http.StatusOK, true, account)
		return
	}
	utils.APIResponse(ctx, "Failed to update user account", http.StatusBadRequest, false, nil)
}

func (user *UsersHandlerImpl) GetAllUsers(ctx *gin.Context) {
	users, err := user.UserService.GetAllUsers(ctx)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if len(*users) > 0 {
		utils.APIResponse(ctx, "User successfully returned", http.StatusOK, true, users)
		return
	}
	utils.APIResponse(ctx, "Failed to return users", http.StatusBadRequest, false, nil)
}

func (user *UsersHandlerImpl) GetUserByPhoneNo(ctx *gin.Context) {
	fetchReq := FetchByPhoneRequest{}
	fetchReq.PhoneNo = ctx.Param("phone_no")
	if err := ctx.ShouldBindJSON(&fetchReq); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}
	userRes, err := user.UserService.GetUserByPhone(ctx, fetchReq.PhoneNo)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if userRes.ID > 0 {
		utils.APIResponse(ctx, "User successfully returned", http.StatusOK, true, userRes)
		return
	}
	utils.APIResponse(ctx, "Failed to return user", http.StatusBadRequest, false, nil)
}

func (user *UsersHandlerImpl) GetUserByEmail(ctx *gin.Context) {
	fetchReq := FetchByIdEmailRequest{}
	fetchReq.Email = ctx.Param("email")
	if err := ctx.ShouldBindJSON(&fetchReq); err != nil {
		utils.APIResponse(ctx, "Failed to convert request to JSON", http.StatusBadRequest,
			false, err.Error())
		return
	}
	userRes, err := user.UserService.GetUserByEmail(ctx, fetchReq.Email)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if userRes.ID > 0 {
		utils.APIResponse(ctx, "User successfully returned", http.StatusOK, true, userRes)
		return
	}
	utils.APIResponse(ctx, "Failed to return user", http.StatusBadRequest, false, nil)
}

func (user *UsersHandlerImpl) GetUserById(ctx *gin.Context) {
	fetchReq := FetchByIdRequest{}
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
	userRes, err := user.UserService.GetUserById(ctx, fetchReq.UserId)
	if err != nil {
		utils.APIResponse(ctx, "Failed, please try again later", http.StatusInternalServerError,
			false, err.Error())
		return
	}
	if userRes.ID > 0 {
		utils.APIResponse(ctx, "User successfully returned", http.StatusOK, true, userRes)
		return
	}
	utils.APIResponse(ctx, "Failed to return user", http.StatusBadRequest, false, nil)
}
