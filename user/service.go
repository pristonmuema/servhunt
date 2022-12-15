package user

import (
	"context"
	"servhunt/infra/token"
	"servhunt/user/dao"
	"strings"
	"time"
)

type UserService interface {
	Login(ctx context.Context, request LoginRequest) (*LoginResponse, error)
	CreateUserAccount(ctx context.Context, user CreateUserRequest) (*CreateUserResponse, error)
	UpdateUserAccount(ctx context.Context, user UpdateUserRequest) (*CreateUserResponse, error)
	GetAllUsers(ctx context.Context) (*[]Response, error)
	GetUserById(ctx context.Context, id int) (*Response, error)
	GetUserByPhone(ctx context.Context, phone string) (*Response, error)
	GetUserByEmail(ctx context.Context, email string) (*Response, error)
}

type UserServiceImpl struct {
	dao.UserRepo
	token.Maker
}

func NewUserServiceImpl(userDao dao.UserRepo, token token.Maker) UserService {
	return &UserServiceImpl{
		UserRepo: userDao,
		Maker:    token,
	}
}

func (u *UserServiceImpl) Login(ctx context.Context, request LoginRequest) (*LoginResponse, error) {

	// get the user
	user, err := u.UserRepo.GetUserByPhone(ctx, request.PhoneNo)
	if err != nil {
		return nil, err
	}
	// check password
	if checkErr := token.CheckPassword(request.Password, user.Password); checkErr != nil {
		return nil, checkErr
	}

	// create token
	createdToken, err := u.Maker.CreateToken(user.PhoneNo, time.Hour*5000)
	if err != nil {
		return nil, err
	}
	var langs []string
	if user.Languages != nil {
		for _, lang := range user.Languages {
			langs = append(langs, lang.Language)
		}
	}
	finalUser := Response{
		ID:            user.ID,
		FirstName:     user.FirstName,
		SecondName:    user.SecondName,
		FullName:      strings.Join([]string{user.FirstName, user.SecondName}, ""),
		Email:         user.Email,
		PhoneNo:       user.PhoneNo,
		UserType:      user.PhoneNo,
		Location:      user.Location,
		Currency:      user.Currency,
		Languages:     langs,
		Description:   user.Description,
		Ratings:       user.Ratings,
		OnlineStatus:  user.OnlineStatus,
		CreatedOn:     user.CreatedOn,
		LastUpdatedOn: user.LastUpdatedOn,
	}
	res := LoginResponse{
		AccessToken: createdToken,
		User:        finalUser,
	}
	return &res, nil
}

func (u *UserServiceImpl) CreateUserAccount(ctx context.Context, user CreateUserRequest) (*CreateUserResponse, error) {

	hashedPassword, errH := token.HashPassword(user.Password)
	if errH != nil {
		return nil, errH
	}
	var langs []dao.Language
	if user.Languages != nil {
		for _, lan := range user.Languages {
			lang := dao.Language{
				Language: lan,
			}
			langs = append(langs, lang)
		}
	}

	request := dao.User{
		FirstName:    user.FirstName,
		SecondName:   user.SecondName,
		Email:        user.Email,
		PhoneNo:      user.PhoneNo,
		UserType:     user.UserType,
		Password:     hashedPassword,
		Currency:     user.Currency,
		Languages:    langs,
		Description:  user.Description,
		Ratings:      user.Ratings,
		OnlineStatus: user.OnlineStatus,
		Location:     user.Location,
		Address:      user.Address,
	}
	savedUser, err := u.UserRepo.SaveUser(ctx, request)
	if err != nil {
		return nil, err
	}

	res := CreateUserResponse{
		UserId: savedUser.ID,
	}
	return &res, nil
}

func (u *UserServiceImpl) UpdateUserAccount(ctx context.Context, user UpdateUserRequest) (*CreateUserResponse, error) {

	var langs []dao.Language
	if user.Languages != nil {
		for _, lan := range user.Languages {
			lang := dao.Language{
				Language: lan,
			}
			langs = append(langs, lang)
		}
	}
	request := dao.User{
		FirstName:     user.FirstName,
		SecondName:    user.SecondName,
		Email:         user.Email,
		PhoneNo:       user.PhoneNo,
		UserType:      user.UserType,
		Currency:      user.Currency,
		Languages:     langs,
		Description:   user.Description,
		Ratings:       user.Ratings,
		OnlineStatus:  user.OnlineStatus,
		Location:      user.Location,
		Address:       user.Address,
		AvailableTime: user.AvailableTime,
		About:         user.About,
	}
	updatedUser, err := u.UserRepo.UpdateUser(ctx, request)
	if err != nil {
		return nil, err
	}
	res := CreateUserResponse{
		UserId: updatedUser.ID,
	}

	return &res, nil
}

func (u *UserServiceImpl) GetAllUsers(ctx context.Context) (*[]Response, error) {

	users, err := u.UserRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var userList []Response
	for _, user := range *users {
		var langs []string
		if user.Languages != nil {
			for _, lang := range user.Languages {
				langs = append(langs, lang.Language)
			}
		}
		finalUser := Response{
			ID:            user.ID,
			FirstName:     user.FirstName,
			SecondName:    user.SecondName,
			FullName:      strings.Join([]string{user.FirstName, user.SecondName}, " "),
			Email:         user.Email,
			PhoneNo:       user.PhoneNo,
			UserType:      user.PhoneNo,
			Location:      user.Location,
			Currency:      user.Currency,
			Languages:     langs,
			Description:   user.Description,
			Ratings:       user.Ratings,
			OnlineStatus:  user.OnlineStatus,
			CreatedOn:     user.CreatedOn,
			LastUpdatedOn: user.LastUpdatedOn,
		}
		userList = append(userList, finalUser)
	}
	return &userList, nil
}

func (u *UserServiceImpl) GetUserById(ctx context.Context, id int) (*Response, error) {
	user, err := u.UserRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	var langs []string
	if user.Languages != nil {
		for _, lang := range user.Languages {
			langs = append(langs, lang.Language)
		}
	}
	finalUser := Response{
		ID:            user.ID,
		FirstName:     user.FirstName,
		SecondName:    user.SecondName,
		FullName:      strings.Join([]string{user.FirstName, user.SecondName}, ""),
		Email:         user.Email,
		PhoneNo:       user.PhoneNo,
		UserType:      user.PhoneNo,
		Location:      user.Location,
		Currency:      user.Currency,
		Languages:     langs,
		Description:   user.Description,
		Ratings:       user.Ratings,
		OnlineStatus:  user.OnlineStatus,
		CreatedOn:     user.CreatedOn,
		LastUpdatedOn: user.LastUpdatedOn,
	}
	return &finalUser, nil
}

func (u *UserServiceImpl) GetUserByPhone(ctx context.Context, phone string) (*Response, error) {
	user, err := u.UserRepo.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	var langs []string
	if user.Languages != nil {
		for _, lang := range user.Languages {
			langs = append(langs, lang.Language)
		}
	}
	finalUser := Response{
		ID:            user.ID,
		FirstName:     user.FirstName,
		SecondName:    user.SecondName,
		FullName:      strings.Join([]string{user.FirstName, user.SecondName}, ""),
		Email:         user.Email,
		PhoneNo:       user.PhoneNo,
		UserType:      user.PhoneNo,
		Location:      user.Location,
		Currency:      user.Currency,
		Languages:     langs,
		Description:   user.Description,
		Ratings:       user.Ratings,
		OnlineStatus:  user.OnlineStatus,
		CreatedOn:     user.CreatedOn,
		LastUpdatedOn: user.LastUpdatedOn,
	}
	return &finalUser, nil
}

func (u *UserServiceImpl) GetUserByEmail(ctx context.Context, email string) (*Response, error) {
	user, err := u.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	var langs []string
	if user.Languages != nil {
		for _, lang := range user.Languages {
			langs = append(langs, lang.Language)
		}
	}
	finalUser := Response{
		ID:            user.ID,
		FirstName:     user.FirstName,
		SecondName:    user.SecondName,
		FullName:      strings.Join([]string{user.FirstName, user.SecondName}, ""),
		Email:         user.Email,
		PhoneNo:       user.PhoneNo,
		UserType:      user.PhoneNo,
		Location:      user.Location,
		Currency:      user.Currency,
		Languages:     langs,
		Description:   user.Description,
		Ratings:       user.Ratings,
		OnlineStatus:  user.OnlineStatus,
		CreatedOn:     user.CreatedOn,
		LastUpdatedOn: user.LastUpdatedOn,
	}
	return &finalUser, nil
}
