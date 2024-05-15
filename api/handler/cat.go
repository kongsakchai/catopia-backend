package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/payload"
	"github.com/kongsakchai/catopia-backend/api/response"
	"github.com/kongsakchai/catopia-backend/domain"
	"github.com/kongsakchai/catopia-backend/utils/mapping"
)

type CatHandler struct {
	catUsecase domain.CatUsecase
}

func NewCatHandler(catUsecase domain.CatUsecase) *CatHandler {
	return &CatHandler{catUsecase}
}

// Get By ID godoc
// @description Get cat by ID
// @tags cat
// @security ApiKeyAuth
// @id CatGetByIDHandler
// @accept json
// @produce json
// @param id path int true "id of cat"
// @response 200 {object} domain.Cat
// @Router /api/cat/{id} [get]
func (h *CatHandler) GetByID(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	data, err := h.catUsecase.GetByID(c, id, userID)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", data)
}

// Get all godoc
// @summary All Cat
// @description Get all cat
// @tags cat
// @security ApiKeyAuth
// @id GetAllCatHandler
// @accept json
// @produce json
// @response 200 {array} domain.Cat
// @Router /api/cat [get]
func (h *CatHandler) GetAll(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	data, err := h.catUsecase.GetByUserID(c, userID)
	if err != nil {
		response.NewError(c, err)
		return
	}

	fmt.Println(data)

	response.New(c, http.StatusOK, "success", data)
}

// Create cat godoc
// @description Create new cat
// @tags cat
// @security ApiKeyAuth
// @id CreateCatHandler
// @accept json
// @produce json
// @param cat body payload.CreateCat true "cat"
// @Router /api/cat/ [post]
func (h *CatHandler) Create(c *gin.Context) {
	var req payload.CreateCat
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, err)
		return
	}

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	data, err := mapping.Mapping[domain.Cat](&req)
	if err != nil {
		response.NewError(c, err)
		return
	}

	err = h.catUsecase.Create(c, userID, data)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusCreated, "success", nil)
}

// Update cat godoc
// @description Update cat by ID
// @tags cat
// @security ApiKeyAuth
// @id UpdateCatHandler
// @accept json
// @produce json
// @param cat body payload.UpdateCat true "cat"
// @param id path int true "id of cat"
// @Router /api/cat/{id} [put]
func (h *CatHandler) Update(c *gin.Context) {
	var req payload.UpdateCat
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, err)
		return
	}

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	data, err := mapping.Mapping[domain.Cat](&req)
	if err != nil {
		response.NewError(c, err)
		return
	}

	err = h.catUsecase.Update(c, id, userID, data)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", nil)
}

// Delete cat godoc
// @description Delete cat by id
// @tags cat
// @security ApiKeyAuth
// @id DeleteCatHandler
// @accept json
// @produce json
// @param id path int true "id of cat"
// @Router /api/cat/{id} [get]
func (h *CatHandler) Delete(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	err = h.catUsecase.Delete(c, id, userID)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", nil)
}
