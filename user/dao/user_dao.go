package dao

import (
	"context"
	"gorm.io/gorm/clause"
	"servhunt/infra/dao"
)

type UserRepo interface {
	SaveUser(ctx context.Context, request User) (*User, error)
	UpdateUser(ctx context.Context, request User) (*User, error)
	GetAllUsers(ctx context.Context) (*[]User, error)
	GetUserById(ctx context.Context, id int) (*User, error)
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type UserRepoImpl struct {
	repo *dao.Repository
}

func NewUserRepoImpl(repo *dao.Repository) UserRepo {
	return &UserRepoImpl{repo: repo}
}

func (u *UserRepoImpl) SaveUser(ctx context.Context, request User) (*User, error) {
	err := u.repo.DB.WithContext(ctx).Model(User{}).Create(&request).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (u *UserRepoImpl) UpdateUser(ctx context.Context, request User) (*User, error) {
	err := u.repo.DB.WithContext(ctx).Model(&User{}).Where("id = ?", request.ID).Updates(User{
		FirstName:     request.FirstName,
		SecondName:    request.SecondName,
		Email:         request.Email,
		PhoneNo:       request.PhoneNo,
		UserType:      request.UserType,
		Currency:      request.Currency,
		Languages:     request.Languages,
		Description:   request.Description,
		Ratings:       request.Ratings,
		AvailableTime: request.AvailableTime,
		About:         request.About,
		OnlineStatus:  request.OnlineStatus,
		LastUpdatedOn: request.LastUpdatedOn,
	}).Error

	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (u *UserRepoImpl) GetAllUsers(ctx context.Context) (*[]User, error) {
	var users []User
	err := u.repo.DB.WithContext(ctx).Model(&User{}).Preload(clause.Associations).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *UserRepoImpl) GetUserById(ctx context.Context, id int) (*User, error) {
	var user User
	err := u.repo.DB.WithContext(ctx).Model(&User{}).Where("id = ?", id).Preload(clause.Associations).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepoImpl) GetUserByPhone(ctx context.Context, phone string) (*User, error) {
	var user User
	err := u.repo.DB.WithContext(ctx).Model(&User{}).Where("phone_no = ?", phone).Preload(clause.Associations).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepoImpl) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := u.repo.DB.WithContext(ctx).Model(&User{}).Where("email = ?", email).Preload(clause.Associations).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
