package controllers

import (
	"backend_jamu/internal/database"
	"backend_jamu/models"
	"backend_jamu/utils"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type ProductController struct{}

// GetAll godoc
// @Summary      Get all products
// @Description  Retrieve a list of all jamu products
// @Tags         products
// @Produce      json
// @Success      200  {object}  utils.APIResponse{data=[]models.Product}
// @Router       /products [get]
func (pc *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	var items []models.Product
	err := database.DB.NewSelect().Model(&items).Scan(context.Background())
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Gagal mengambil data produk")
		return
	}
	utils.JSONResponse(w, http.StatusOK, "Daftar produk berhasil diambil", items)
}

// Create godoc
// @Summary      Create a new product
// @Description  Add a new product to the catalog (Admin Only)
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      models.Product  true  "Product Data"
// @Success      201      {object}  utils.APIResponse{data=models.Product}
// @Security     BearerAuth
// @Router       /products [post]
func (pc *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	item := new(models.Product)
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Payload tidak valid")
		return
	}

	_, err := database.DB.NewInsert().Model(item).Exec(context.Background())
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Gagal menambahkan produk")
		return
	}
	utils.JSONResponse(w, http.StatusCreated, "Produk berhasil ditambahkan", item)
}

// Update godoc
// @Summary      Update a product
// @Description  Update product details by ID (Admin Only)
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       query     int             true  "Product ID"
// @Param        product  body      models.Product  true  "Product Data"
// @Success      200      {object}  utils.APIResponse{data=models.Product}
// @Security     BearerAuth
// @Router       /products [put]
func (pc *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	item := new(models.Product)
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Payload tidak valid")
		return
	}
	item.ID = id

	res, err := database.DB.NewUpdate().
		Model(item).
		Column("name", "description", "price", "image").
		WherePK().
		Exec(context.Background())

	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Gagal memperbarui produk")
		return
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		utils.JSONError(w, http.StatusNotFound, "Produk tidak ditemukan")
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Produk berhasil diperbarui", item)
}

// Delete godoc
// @Summary      Delete a product
// @Description  Remove a product from the catalog by ID (Admin Only)
// @Tags         products
// @Param        id   query     int  true  "Product ID"
// @Success      200  {object}  utils.APIResponse
// @Security     BearerAuth
// @Router       /products [delete]
func (pc *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	res, err := database.DB.NewDelete().
		Model((*models.Product)(nil)).
		Where("id = ?", id).
		Exec(context.Background())

	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Gagal menghapus produk")
		return
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		utils.JSONError(w, http.StatusNotFound, "Produk tidak ditemukan")
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Produk berhasil dihapus", nil)
}
