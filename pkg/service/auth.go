package service

import (
	"strconv"
	"tasktracker-api/pkg/models"
	"tasktracker-api/pkg/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	usersRepo repository.Users
	tasksRepo repository.Tasks
}

type tokenClaims struct {
	jwt.RegisteredClaims
}

const (
	TokenLifeTime = 24 * time.Hour
)

var jwtSignedKey = []byte("secret")

func NewAuthService(usersRepo repository.Users, tasksRepo repository.Tasks) *AuthService {
	return &AuthService{
		usersRepo: usersRepo,
		tasksRepo: tasksRepo,
	}
}

func (s *AuthService) Login(user models.AuthData) (string, error) {
	var jwtToken string
	userId, err := s.CheckUser(*user.Username, *user.Password)
	if err != nil {
		return jwtToken, err
	}
	jwtToken, err = s.GenerateToken(userId)
	if err != nil {
		return jwtToken, err
	}
	return jwtToken, nil
}

func (s *AuthService) Logup(user models.UserData) (string, error) {
	var jwtToken string
	userId, err := s.CreateUser(user)
	if err != nil {
		return jwtToken, err
	}
	jwtToken, err = s.GenerateToken(userId)
	if err != nil {
		return jwtToken, err
	}
	return jwtToken, nil
}

func (s *AuthService) GenerateToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenLifeTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.Itoa(id),
		},
	})
	tokenString, err := token.SignedString(jwtSignedKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *AuthService) GetUserById(userId int) (models.User, error) {
	return s.usersRepo.GetUserById(userId)
}

func (s *AuthService) CheckUser(username string, password string) (int, error) {
	userFromDB, err := s.usersRepo.GetUserByLogin(username)
	if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(password))
	if err != nil {
		return 0, err
	}
	return userFromDB.Id, nil
}

func (s *AuthService) CreateUser(user models.UserData) (int, error) {
	Password, err := s.HashedPassword(*user.Password)
	user.Password = &Password
	if err != nil {
		return 0, err
	}
	return s.usersRepo.CreateUser(user)
}

func (s *AuthService) HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
