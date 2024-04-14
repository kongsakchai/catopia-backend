package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/payload"
	"github.com/kongsakchai/catopia-backend/api/response"
	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
	"github.com/kongsakchai/catopia-backend/utils/data"
)

type CatHandler struct {
	catUsecase domain.CatUsecase
}

func NewCatHandler(catUsecase domain.CatUsecase) *CatHandler {
	return &CatHandler{catUsecase}
}

func (h *CatHandler) GetByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	data, err := h.catUsecase.GetByID(c, id, userID)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusOK, "success", data)
}

func (h *CatHandler) GetAll(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	data, err := h.catUsecase.GetByUserID(c, userID)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusOK, "success", data)
}

func (h *CatHandler) Create(c *gin.Context) {
	var req payload.CreateCat
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	data, err := data.Mapping[domain.CatModel](&req)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	err = h.catUsecase.Create(c, data, userID)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusCreated, "success", nil)
}

func (h *CatHandler) Update(c *gin.Context) {
	var req payload.UpdateCat
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	data, err := data.Mapping[domain.CatModel](&req)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	err = h.catUsecase.Update(c, id, userID, data)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusOK, "success", nil)
}

func (h *CatHandler) Delete(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	err = h.catUsecase.Delete(c, id, userID)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusOK, "success", nil)
}
