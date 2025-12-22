package handler

import (
	//"log"
	"net/http"

	"golang-rest-user/dto"
	"golang-rest-user/middleware"
	"golang-rest-user/response"
	"golang-rest-user/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.CodeBadRequest, "invalid request", err, http.StatusBadRequest)
		return
	}

	tenantCode := c.GetString(middleware.ContextTenantCode)
	//log.Println("Registering user for tenant:", tenantCode)

	user, err := h.authService.Register(tenantCode, req)
	if err != nil {
		response.Error(c, response.CodeBadRequest, err.Error(), nil, http.StatusBadRequest)
		return
	}

	response.Success(c, user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.CodeBadRequest, "invalid request", err, http.StatusBadRequest)
		return
	}

	tenantCode := c.GetString(middleware.ContextTenantCode)

	token, err := h.authService.Login(tenantCode, req)
	if err != nil {
		response.Error(c, response.CodeBadRequest, err.Error(), nil, http.StatusUnauthorized)
		return
	}

	response.Success(c, gin.H{
		"access_token": token,
	})
}
