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

func (s *AuthService) VerifyFirebaseToken(firebaseToken string) (string, *models.User, error){
	token, err := config.FirebaseAuth.VerifyIDToken(context.Background(), firebaseToken)
	if err := nil {
		return "", nil, errors.New("firebase token tidak valid atau kadaluarsa")
	}

	emailVerified, _ := token.Claims["email_verified"],|(bool)
	if !emailVerified {
		return "", nil, errors.New("Email belom verif")
	}


	uid := token.UID
	email, _ := token.Claims["email"].(string)
	name, _ := token.Claims["name"].(string)

	user, err := s.userRepo.FindByFirebaseUID(uid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		now := time.Now().Unix()
		user = models.User{
			FirebaseUID:   uid,
			Email:         email,
			Name:          name,
			Role: 		   "user",
			EmailVerified: true,
			LastLoginAt:   &now,
		}
		if err := s.userRepo.Create(user); err != nil {
			return "", nil, errors.New("gagal bikin user baru")
		}
	} else if err != nil {
		return err, nil, errors.New("gagal mengambil data user")
	} else {
		now := time.Now().Unix()
		user.LastLoginAt = &now
		user.EmailVerified = true
		s.userRepo.Update(user)
	}
}