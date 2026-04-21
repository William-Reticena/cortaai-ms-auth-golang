package request

type LoginRequest struct {
	DsEmail    string `json:"dsEmail" binding:"required,email"`
	DsPassword string `json:"dsPassword" binding:"required,min=6"`
}

type RegisterRequest struct {
	DsEmail    string `json:"dsEmail" binding:"required,email"`
	DsPassword string `json:"dsPassword" binding:"required,min=6"`
	NmUser     string `json:"nmUser" binding:"required,min=3"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type ValidateTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

