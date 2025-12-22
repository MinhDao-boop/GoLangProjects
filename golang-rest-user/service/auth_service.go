package service

import (
	"errors"

	"golang-rest-user/dto"
	"golang-rest-user/models"
	"golang-rest-user/repository"
	"golang-rest-user/security"
)

type AuthService interface {
	Register(tenantCode string, req dto.CreateUserRequest) (*models.User, error)
	Login(tenantCode string, req dto.LoginRequest) (string, error)
}

type authService struct {
	userRepoFactory func(string) (repository.UserRepo, error)
	jwtManager      *security.Manager
}

func NewAuthService(
	userRepoFactory func(string) (repository.UserRepo, error),
	jwtManager *security.Manager,
) AuthService {
	return &authService{
		userRepoFactory: userRepoFactory,
		jwtManager:      jwtManager,
	}
}

func (s *authService) Register(tenantCode string, req dto.CreateUserRequest) (*models.User, error) {

	repo, err := s.userRepoFactory(tenantCode)
	if err != nil {
		return nil, err
	}

	if _, err := repo.GetByUsername(req.Username); err == nil {
		return nil, errors.New("username already exists")
	}

	encryptedPass, err := security.Encrypt(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Password: encryptedPass,
		FullName: req.FullName,
	}

	if err := repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(tenantCode string, req dto.LoginRequest) (string, error) {

	repo, err := s.userRepoFactory(tenantCode)
	if err != nil {
		return "", err
	}

	user, err := repo.GetByUsername(req.Username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	decryptedPass, _ := security.Decrypt(user.Password)
	if decryptedPass != req.Password {
		return "", errors.New("invalid credentials")
	}

	return s.jwtManager.GenerateAccessToken(
		user.ID,
		user.Username,
		tenantCode,
	)
}
