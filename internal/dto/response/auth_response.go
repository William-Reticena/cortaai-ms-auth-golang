package response

type LoginResponse struct {
	AccessToken  string        `json:"dsAccessToken"`
	RefreshToken string        `json:"dsRefreshToken"`
	TokenType    string        `json:"dsTokenType"`
	ExpiresIn    int           `json:"nrExpiresIn"`
	User         *UserResponse `json:"user"`
}

type RegisterResponse struct {
	Message string        `json:"message"`
	User    *UserResponse `json:"user"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"dsAccessToken"`
	TokenType   string `json:"dsTokenType"`
	ExpiresIn   int    `json:"nrExpiresIn"`
}

type ValidateTokenResponse struct {
	InValid bool          `json:"inValid"`
	User    *UserResponse `json:"user,omitempty"`
}

type UserResponse struct {
	ID      int    `json:"id"`
	DsEmail string `json:"dsEmail"`
	NmUser  string `json:"nmUser"`
}
