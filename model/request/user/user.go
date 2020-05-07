package user

// CreateUser Struct
type CreateUser struct {
	Name     string `json:"name,omitempty"`
	UserName string `json:"userName,omitempty"`
	PassWord string `json:"passWord,omitempty"`
	Role     string `json:"role,omitempty"`
}

// UpdateUser Struct
type UpdateUser struct {
	ParamUserID string
	Name        *string `json:"name,omitempty"`
	UserName    *string `json:"userName,omitempty"`
	PassWord    *string `json:"passWord,omitempty"`
	Role        *string `json:"role,omitempty"`
}

// SignInUser Struct
type SignInUser struct {
	UserName string `json:"userName,omitempty"`
	PassWord string `json:"passWord,omitempty"`
}
