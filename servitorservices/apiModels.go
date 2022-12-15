package servitorservices

type ServiceRequest struct {
	UserID          int         `json:"user_id"`
	ServiceImage    string      `json:"service_image"`
	ServiceName     string      `json:"service_name"`
	ServiceDuration string      `json:"service_duration"`
	ServiceCost     float64     `json:"service_cost"`
	Locations       []Locations `json:"locations"`
	Category        []Category  `json:"categories"`
}

type ServicesResponse struct {
	ID              int         `json:"id"`
	UserID          int         `json:"user_id"`
	ServiceImage    string      `json:"service_image"`
	ServiceName     string      `json:"service_name"`
	ServiceDuration string      `json:"service_duration"`
	ServiceCost     float64     `json:"service_cost"`
	Locations       []Locations `json:"locations"`
	Category        []Category  `json:"categories"`
}

type UpdateServiceRequest struct {
	ID              int     `json:"id"`
	ServiceImage    string  `json:"service_image"`
	ServiceName     string  `json:"service_name"`
	ServiceDuration string  `json:"service_duration"`
	ServiceCost     float64 `json:"service_cost"`
}

type Locations struct {
	LocationImage string  `json:"location_image"`
	LocationName  string  `json:"location_name"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Address       string  `json:"address"`
}

type Category struct {
	CategoryImage string `json:"category_image"`
	CategoryName  string `json:"category_name"`
}

type ServiceResponse struct {
	ServiceId int `json:"service_id"`
}

type FetchByIdRequest struct {
	ServiceId int `json:"service_id"`
}

type OrderIdRequest struct {
	UserId int `json:"user_id"`
}

type LocationInfoRequest struct {
	LocationImage string  `json:"location_image"`
	LocationName  string  `json:"location_name"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Address       string  `json:"address"`
	ServiceID     int     `json:"service_id"`
}

type UpdateLocationInfoRequest struct {
	ID            int     `json:"id"`
	LocationImage string  `json:"location_image"`
	LocationName  string  `json:"location_name"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Address       string  `json:"address"`
	ServiceID     int     `json:"service_id"`
}

type LocationInfoResponse struct {
	LocationID int `json:"location_id"`
}

type FetchLocByIdRequest struct {
	LocationID int `json:"id"`
}

type CategoryRequest struct {
	CategoryImage string `json:"category_image"`
	CategoryName  string `json:"category_name"`
	ServiceID     int    `json:"service_id"`
}

type UpdateCategoryRequest struct {
	CategoryImage string `json:"category_image"`
	CategoryName  string `json:"category_name"`
	ServiceID     int    `json:"service_id"`
}

type CategoryResponse struct {
	CategoryId int `json:"category_id"`
}

type FetchCatByIdRequest struct {
	CategoryId int `json:"id"`
}

type LocationResponse struct {
	ID            int     `json:"id"`
	ServiceID     int     `json:"service_id"`
	LocationImage string  `json:"location_image"`
	LocationName  string  `json:"location_name"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Address       string  `json:"address"`
}
type CategoriesResponse struct {
	ID            int    `json:"id"`
	ServiceID     int    `json:"service_id"`
	CategoryImage string `json:"category_image"`
	CategoryName  string `json:"category_name"`
}

type ChanResponse struct {
	Data interface{}
	Err  error
}

type ServHuntChan chan ChanResponse
