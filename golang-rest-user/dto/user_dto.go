package dto

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,email"`
	FullName string `json:"fullname" binding:"required"`
	Phone    string `json:"phone" binding:"omitempty"`
	Position string `json:"position" binding:"omitempty"`
}

type UpdateUserRequest struct {
	FullName string `json:"fullname" binding:"required"`
	Phone    string `json:"phone" binding:"omitempty"`
	Position string `json:"position" binding:"omitempty"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FullName  string `json:"fullname"`
	Phone     string `json:"phone"`
	Position  string `json:"position"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListUsersResponse struct {
	Data     []UserResponse `json:"data"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
	Total    int64          `json:"total"`
}
