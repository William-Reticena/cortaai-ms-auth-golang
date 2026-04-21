package main

import (
	"log"
	"os"
	"path/filepath"

	"cortaai-ms-auth-go/internal/controller"
	"cortaai-ms-auth-go/internal/database"
	"cortaai-ms-auth-go/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	envPath := filepath.Join(os.Getenv("PWD"), ".env")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		envPath = filepath.Join(os.Getenv("PWD"), "..", ".env")
	}

	if _, err := os.Stat(envPath); err == nil {
		godotenv.Load(envPath)
		log.Printf("✅ Carregado .env de: %s\n", envPath)
	}
}

func main() {
	if os.Getenv("JWT_SECRET") == "" {
		os.Setenv("JWT_SECRET", "seu-secret-super-seguro-aqui-mude-em-producao")
	}

	if err := database.InitDB(); err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco: %v", err)
	}

	server := gin.Default()

	authService := &services.AuthService{}
	authController := &controller.AuthController{
		AuthService: authService,
	}

	auth := server.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.POST("/refresh", authController.RefreshToken)
		auth.POST("/validate", authController.ValidateToken)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	log.Printf("🚀 Servidor iniciado em http://localhost:%s\n", port)
	server.Run(":" + port)
}
