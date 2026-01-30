package controllers

import (
	"backend_jamu/internal/database"
	"backend_jamu/models"
	"backend_jamu/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type NewsController struct{}

// GetAll godoc
// @Summary      Get all news
// @Description  Retrieve a list of all news articles ordered by date
// @Tags         news
// @Produce      json
// @Success      200  {object}  utils.APIResponse{data=[]models.News}
// @Router       /news [get]
func (nc *NewsController) GetAll(w http.ResponseWriter, r *http.Request) {
	var items []models.News
	err := database.DB.NewSelect().
		Model(&items).
		Order("created_at DESC").
		Scan(context.Background())

	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Gagal mengambil data berita")
		return
	}
	utils.JSONResponse(w, http.StatusOK, "Daftar berita berhasil diambil", items)
}

// Create godoc
// @Summary      Create news
// @Description  Add a new news article (Admin Only)
// @Tags         news
// @Accept       json
// @Produce      json
// @Param        news  body      models.News  true  "News Data"
// @Success      201   {object}  utils.APIResponse{data=models.News}
// @Security     BearerAuth
// @Router       /news [post]
func (nc *NewsController) Create(w http.ResponseWriter, r *http.Request) {
	item := new(models.News)
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Payload tidak valid")
		return
	}

	_, err := database.DB.NewInsert().Model(item).Exec(context.Background())
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Gagal menambahkan berita")
		return
	}
	utils.JSONResponse(w, http.StatusCreated, "Berita berhasil ditambahkan", item)
}

// Update godoc
// @Summary      Update news
// @Description  Update news article by ID (Admin Only)
// @Tags         news
// @Accept       json
// @Produce      json
// @Param        id    query     int          true  "News ID"
// @Param        news  body      models.News  true  "News Data"
// @Success      200   {object}  utils.APIResponse{data=models.News}
// @Security     BearerAuth
// @Router       /news [put]
func (nc *NewsController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	fmt.Printf("[NewsController] Update Request for ID: %s\n", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	item := new(models.News)
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Payload tidak valid")
		return
	}
	item.ID = id

	res, err := database.DB.NewUpdate().
		Model(item).
		Column("title", "content").
		WherePK().
		Exec(context.Background())

	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Gagal memperbarui berita")
		return
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		utils.JSONError(w, http.StatusNotFound, "Berita tidak ditemukan")
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Berita berhasil diperbarui", item)
}

// Delete godoc
// @Summary      Delete news
// @Description  Remove a news article by ID (Admin Only)
// @Tags         news
// @Param        id   query     int  true  "News ID"
// @Success      200  {object}  utils.APIResponse
// @Security     BearerAuth
// @Router       /news [delete]
func (nc *NewsController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	fmt.Printf("[NewsController] Delete Request for ID: %s\n", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	res, err := database.DB.NewDelete().
		Model((*models.News)(nil)).
		Where("id = ?", id).
		Exec(context.Background())

	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Gagal menghapus berita")
		return
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		utils.JSONError(w, http.StatusNotFound, "Berita tidak ditemukan")
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Berita berhasil dihapus", nil)
}
