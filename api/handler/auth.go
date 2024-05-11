package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/payload"
	"github.com/kongsakchai/catopia-backend/api/response"
	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
	"github.com/kongsakchai/catopia-backend/utils/mapping"
)

type AuthHandler struct {
	authUsecase domain.AuthUsecase
}

func NewAuthHandler(authUsecase domain.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase}
}

// HealthCheckHandler godoc
// @summary Health Check
// @description Health checking for the service
// @id HealthCheckHandler
// @produce plain
// @response 200 {string} string "OK"
// @router /healthcheck [get]
func (h *AuthHandler) HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// Register godoc
// @description Register new user
// @tags auth
// @security ApiKeyAuth
// @id RegisterHandler
// @accept json
// @produce json
// @param User body payload.Regis true "User data"
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req payload.Regis
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, err)
		return
	}

	data, err := mapping.Mapping[domain.User](&req)
	if err != nil {
		response.NewError(c, err)
		return
	}

	data.Password = req.Password
	err = h.authUsecase.Register(c, data)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusCreated, "success", nil)
}

// Login godoc
// @description Login and get token
// @tags auth
// @security ApiKeyAuth
// @id LoginHandler
// @accept json
// @produce json
// @param User body payload.Login true "User data"
// @response 201 {object} payload.LoginResponse
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req payload.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, err)
		return
	}

	first, token, err := h.authUsecase.Login(c, req.Username, req.Password)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusCreated, "success", &payload.LoginResponse{FirstLogin: first, Token: token})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	id := c.Param("session_id")
	if id == "" {
		response.NewError(c, errs.NewError(errs.ErrBadRequest, nil))
		return
	}

	err := h.authUsecase.Logout(c, id)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", nil)
}
