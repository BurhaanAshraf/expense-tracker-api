package handler

import (
	"encoding/json"
	"net/http"

	"github.com/BurhaanAshraf/finance-api/internal/dto"
	"github.com/BurhaanAshraf/finance-api/internal/model"
	"github.com/BurhaanAshraf/finance-api/internal/response"
	service "github.com/BurhaanAshraf/finance-api/internal/service"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	user, err := h.userService.Register(
		r.Context(),
		req.Name,
		req.Email,
		req.Password,
	)

	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	userResponse := model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	response.Created(w, userResponse)
}
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.BadRequest(w, "Invalid request body")
		return
	}

	token, err := h.userService.Login(
		r.Context(),
		req.Email,
		req.Password,
	)

	if err != nil {
		response.Unauthorized(w, err.Error())
		return
	}

	response.OK(w, dto.LoginResponse{
		Token: token,
	})
}
