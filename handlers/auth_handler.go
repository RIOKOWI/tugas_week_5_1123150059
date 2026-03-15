package handlers

import (
	"github.com/RIOKOWI/tugas_week_5_1123150059/services"
)


type AuthHandler struct {
	authService *services.AuthService
}
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{authService: services.NewAuthService()}
}