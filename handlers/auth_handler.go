package handlers

import (
	"net/http"

	"github.com/RIOKOWI/tugas_week_5_1123150059/services"
	"github.com/gin-gonic/gin"
)


type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{authService: services.NewAuthService()}
}

func (h *AuthHandler) VerifyToken(c *gin.Context) {
	var req struct {
		FirebaseToken string `json:"firebase_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "firebase_token wajib diisi",
		})
		return
	}
}