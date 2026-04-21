package response

type LoginResponse struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	TokenType    string        `json:"token_type"`
	ExpiresIn    int           `json:"expires_in"`
	User         *UserResponse `json:"user"`
}

type RegisterResponse struct {
	Message string        `json:"message"`
	User    *UserResponse `json:"user"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type ValidateTokenResponse struct {
	Valid bool          `json:"valid"`
	User  *UserResponse `json:"user,omitempty"`
}

type UserResponse struct {
	ID      int    `json:"id"`
	DsEmail string `json:"dsEmail"`
	NmUser  string `json:"nmUser"`
}

