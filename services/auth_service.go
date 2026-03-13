package services

import (
	"context"
	"errors"
	"cs"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/RIOKOWI/tugas_week_5_1123150059/config"
	"github.com/RIOKOWI/tugas_week_5_1123150059/models"
	"github.com/RIOKOWI/tugas_week_5_1123150059/repositories"
	"gorm.io/gorm"

)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService(userRepo: repositpries.NewUserRepository())
}