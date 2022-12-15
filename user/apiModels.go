package user

import "time"

type LoginRequest struct {
	PhoneNo  string `json:"phone_no" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	AccessToken string   `json:"access_token"`
	User        Response `json:"user"`
}

type CreateUserRequest struct {
	FirstName     string   `json:"first_name" binding:"required,alphanum"`
	SecondName    string   `json:"second_name" binding:"required,alphanum"`
	Email         string   `json:"email" binding:"required,email"`
	PhoneNo       string   `json:"phone_no" binding:"required"`
	UserType      string   `json:"user_type" binding:"required"`
	Password      string   `json:"password" binding:"required,min=6"`
	Location      string   `json:"location"`
	Address       string   `json:"address"`
	Currency      string   `json:"currency"`
	Languages     []string `json:"languages"`
	Description   string   `json:"description"`
	Ratings       string   `json:"ratings"`
	OnlineStatus  string   `json:"online_status"`
	AvailableTime string   `json:"available_time"`
	About         string   `json:"about"`
}

type UpdateUserRequest struct {
	ID            int      `json:"id" binding:"required"`
	FirstName     string   `json:"first_name"`
	SecondName    string   `json:"second_name"`
	Email         string   `json:"email"`
	PhoneNo       string   `json:"phone_no"`
	UserType      string   `json:"user_type"`
	Location      string   `json:"location"`
	Address       string   `json:"address"`
	Currency      string   `json:"currency"`
	Languages     []string `json:"languages"`
	Description   string   `json:"description"`
	Ratings       string   `json:"ratings"`
	OnlineStatus  string   `json:"online_status"`
	AvailableTime string   `json:"available_time"`
	About         string   `json:"about"`
}

type CreateUserResponse struct {
	UserId int `json:"user_id"`
}

type FetchByIdRequest struct {
	UserId int `json:"user_id"`
}

type FetchByPhoneRequest struct {
	PhoneNo string `json:"phone_no"`
}

type FetchByIdEmailRequest struct {
	Email string `json:"email" `
}

type Response struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"first_name"`
	SecondName    string    `json:"second_name"`
	FullName      string    `json:"full_name"`
	Email         string    `json:"email"`
	PhoneNo       string    `json:"phone_no"`
	UserType      string    `json:"user_type"`
	Location      string    `json:"location"`
	Currency      string    `json:"currency"`
	Languages     []string  `json:"languages"`
	Description   string    `json:"description"`
	Ratings       string    `json:"ratings"`
	OnlineStatus  string    `json:"online_status"`
	CreatedOn     time.Time `json:"created_on"`
	LastUpdatedOn time.Time `json:"last_updated_on"`
}

type LocationInfoResponse struct {
	ID            int       `json:"id"`
	LocationImage string    `json:"location_image"`
	LocationName  string    `json:"location_name"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	Address       string    `json:"address"`
	UserID        int       `json:"user_id"`
	CreatedOn     time.Time `json:"created_on"`
	LastUpdatedOn time.Time `json:"last_updated_on"`
}
