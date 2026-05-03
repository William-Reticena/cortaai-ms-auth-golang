package controller

import (
	"net/http"

	"cortaai-ms-auth-go/internal/dto"
	"cortaai-ms-auth-go/internal/dto/request"
	"cortaai-ms-auth-go/internal/dto/response"
	"cortaai-ms-auth-go/internal/services"
	"cortaai-ms-auth-go/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
}

func (a *AuthController) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:       "invalid_request",
			Description: err.Error(),
		})
		return
	}

	user, err := a.AuthService.LoginUser(req.DsEmail, req.DsPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:       "invalid_credentials",
			Description: err.Error(),
		})
		return
	}

	accessToken, err := utils.GenerateToken(user.ID, user.Email, utils.GetTokenExpirationMinutes())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:       "token_generation_failed",
			Description: err.Error(),
		})
		return
	}

	refreshToken, err := utils.GenerateToken(user.ID, user.Email, 7*24*60)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:       "token_generation_failed",
			Description: err.Error(),
		})
		return
	}

	expiresIn := utils.GetTokenExpirationMinutes() * 60
	c.JSON(http.StatusOK, response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    expiresIn,
		User: &response.UserResponse{
			ID:      user.ID,
			DsEmail: user.Email,
			NmUser:  user.Username,
		},
	})
}

func (a *AuthController) Register(c *gin.Context) {
	var req request.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:       "invalid_request",
			Description: err.Error(),
		})
		return
	}

	user, err := a.AuthService.RegisterUser(services.RegisterInput{
		DsEmail:    req.DsEmail,
		DsPassword: req.DsPassword,
		NmUser:     req.NmUser,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:       "registration_failed",
			Description: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.RegisterResponse{
		Message: "Usuário criado com sucesso",
		User: &response.UserResponse{
			ID:      user.ID,
			DsEmail: user.Email,
			NmUser:  user.Username,
		},
	})
}

func (a *AuthController) RefreshToken(c *gin.Context) {
	var req request.RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:       "invalid_request",
			Description: err.Error(),
		})
		return
	}

	claims, err := utils.ValidateToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error:       "invalid_token",
			Description: "Token inválido ou expirado",
		})
		return
	}

	newAccessToken, err := utils.GenerateToken(claims.UserID, claims.Email, utils.GetTokenExpirationMinutes())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:       "token_generation_failed",
			Description: err.Error(),
		})
		return
	}

	expiresIn := utils.GetTokenExpirationMinutes() * 60
	c.JSON(http.StatusOK, response.RefreshTokenResponse{
		AccessToken: newAccessToken,
		TokenType:   "Bearer",
		ExpiresIn:   expiresIn,
	})
}

func (a *AuthController) ValidateToken(c *gin.Context) {
	var req request.ValidateTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:       "invalid_request",
			Description: err.Error(),
		})
		return
	}

	claims, err := utils.ValidateToken(req.DsToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ValidateTokenResponse{
			InValid: false,
		})
		return
	}

	user, err := a.AuthService.GetUserByID(claims.UserID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ValidateTokenResponse{
			InValid: false,
		})
		return
	}

	c.JSON(http.StatusOK, response.ValidateTokenResponse{
		InValid: true,
		User: &response.UserResponse{
			ID:      user.ID,
			DsEmail: user.Email,
			NmUser:  user.Username,
		},
	})
}
