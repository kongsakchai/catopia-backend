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

type TreatmentHandler struct {
	treatmentUsecase domain.TreatmentUsecase
}

func NewTreatmentHandler(t domain.TreatmentUsecase) *TreatmentHandler {
	return &TreatmentHandler{
		treatmentUsecase: t,
	}
}

func (t *TreatmentHandler) GetType(c *gin.Context) {
	data, err := t.treatmentUsecase.GetType(c)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusOK, "success", data)
}

func (t *TreatmentHandler) GetByID(c *gin.Context) {
	catID, err := strconv.Atoi(c.Param("cat_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	data, err := t.treatmentUsecase.GetByID(c, id, catID)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusOK, "success", data)
}

func (t *TreatmentHandler) GetByCatID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	catID, err := strconv.Atoi(c.Param("cat_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	data, err := t.treatmentUsecase.GetByCatID(c, catID, userID)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusOK, "success", data)
}

func (t *TreatmentHandler) Create(c *gin.Context) {
	var req payload.CreateTreatment
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	catID, err := strconv.Atoi(c.Param("cat_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	data, err := data.Mapping[domain.TreatmentModel](&req)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	err = t.treatmentUsecase.Create(c, catID, userID, data)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusCreated, "success", nil)
}

func (t *TreatmentHandler) Update(c *gin.Context) {
	var req payload.UpdateTreatment
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	catID, err := strconv.Atoi(c.Param("cat_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	data, err := data.Mapping[domain.TreatmentModel](&req)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	err = t.treatmentUsecase.Update(c, id, catID, data)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusOK, "success", nil)
}

func (t *TreatmentHandler) Delete(c *gin.Context) {
	catID, err := strconv.Atoi(c.Param("cat_id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	err = t.treatmentUsecase.Delete(c, id, catID)
	if err != nil {
		response.NewErrorResponse(c, err)
		return
	}

	response.NewResponse(c, http.StatusOK, "success", nil)
}
