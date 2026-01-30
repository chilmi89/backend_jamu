package controllers

import (
	"backend_jamu/internal/database"
	"backend_jamu/models"
	"backend_jamu/utils"
	"context"
	"encoding/json"
	"net/http"
)

type AuthController struct{}

// Login godoc
// @Summary      User Login
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login  body      models.LoginRequest  true  "Login Credentials"
// @Success      200    {object}  utils.APIResponse{data=models.LoginResponse}
// @Failure      400    {object}  utils.APIResponse
// @Failure      401    {object}  utils.APIResponse
// @Router       /login [post]
func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.JSONError(w, http.StatusMethodNotAllowed, "Metode HTTP tidak diizinkan")
		return
	}

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Payload tidak valid")
		return
	}

	user := new(models.User)
	err := database.DB.NewSelect().
		Model(user).
		Where("email = ?", req.Email).
		Scan(context.Background())

	if err != nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		utils.JSONError(w, http.StatusUnauthorized, "Email atau password salah")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Gagal membuat token akses")
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Login berhasil", models.LoginResponse{
		Token: token,
		User:  *user,
	})
}
