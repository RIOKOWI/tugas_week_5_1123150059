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

	jwtToken, user, err := h.authService.VerifyFirebaseToken(req.FirebaseToken)
	if err != nil {
		if err.Error() == "EMAIL_NOT_VERIFIED" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Email belum diverifikasi. Cek inbox email Anda.",
				"error_code": "EMAIL_NOT_VERIFIED",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": err.Error(),
				"error_code": "INVALID_FIREBASE_TOKEN",
			})
		}
		return
	}
}