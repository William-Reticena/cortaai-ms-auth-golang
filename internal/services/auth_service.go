package services

import (
	"errors"
	"fmt"

	"cortaai-ms-auth-go/internal/database"
	"cortaai-ms-auth-go/internal/model"
	"cortaai-ms-auth-go/internal/utils"
)

type AuthService struct{}

type RegisterInput struct {
	DsEmail    string
	DsPassword string
	NmUser     string
}

func (a *AuthService) RegisterUser(input RegisterInput) (*model.User, error) {
	if input.DsEmail == "" || input.DsPassword == "" {
		return nil, errors.New("email e password são obrigatórios")
	}

	if len(input.DsPassword) < 6 {
		return nil, errors.New("password deve ter pelo menos 6 caracteres")
	}

	var existingUser model.User
	if err := database.DB.Where("ds_email = ?", input.DsEmail).First(&existingUser).Error; err == nil {
		return nil, errors.New("email já cadastrado")
	}

	hashedPassword, err := utils.HashPassword(input.DsPassword)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer hash da senha: %v", err)
	}

	user := &model.User{
		Email:        input.DsEmail,
		PasswordHash: hashedPassword,
		Username:     input.NmUser,
	}

	if err := database.DB.Create(user).Error; err != nil {
		return nil, fmt.Errorf("erro ao criar usuário: %v", err)
	}

	return user, nil
}

func (a *AuthService) LoginUser(email, password string) (*model.User, error) {
	var user model.User

	if err := database.DB.Where("ds_email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("email ou senha incorretos")
	}

	if err := utils.CheckPassword(user.PasswordHash, password); err != nil {
		return nil, errors.New("email ou senha incorretos")
	}

	return &user, nil
}

func (a *AuthService) GetUserByID(userID int) (*model.User, error) {
	var user model.User

	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	return &user, nil
}
