package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/payload"
	"github.com/kongsakchai/catopia-backend/api/response"
	"github.com/kongsakchai/catopia-backend/domain"
	"github.com/kongsakchai/catopia-backend/utils/mapping"
)

type TreatmentHandler struct {
	treatmentUsecase domain.TreatmentUsecase
}

func NewTreatmentHandler(t domain.TreatmentUsecase) *TreatmentHandler {
	return &TreatmentHandler{
		treatmentUsecase: t,
	}
}

// GetType godoc
// @description Get treatment type
// @tags  treatment
// @security ApiKeyAuth
// @id TreatmentGetTypeHandler
// @accept json
// @produce json
// @response 200 {object} domain.TreatmentType
// @Router /api/treatment/type [get]
func (t *TreatmentHandler) GetType(c *gin.Context) {
	data, err := t.treatmentUsecase.GetType(c)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", data)
}

// GetByID godoc
// @description Get treatment by ID
// @tags  treatment
// @security ApiKeyAuth
// @id TreatmentGetByIDHandler
// @accept json
// @produce json
// @param cat_id path int true "cat id"
// @param id path int true "id of treatment"
// @response 200 {object} domain.Treatment
// @Router /api/treatment/{cat_id}/{id} [get]
func (t *TreatmentHandler) GetByID(c *gin.Context) {
	catID, err := strconv.ParseInt(c.Param("cat_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	data, err := t.treatmentUsecase.GetByID(c, id, catID)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", data)
}

// GetByCatID godoc
// @description Get treatment by cat ID
// @tags  treatment
// @security ApiKeyAuth
// @id TreatmentGetByCatIDHandler
// @accept json
// @produce json
// @param cat_id path int true "cat id"
// @response 200 {array} domain.Treatment
// @Router /api/treatment/:user_id/{cat_id} [get]
func (t *TreatmentHandler) GetByCatID(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	catID, err := strconv.ParseInt(c.Param("cat_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	data, err := t.treatmentUsecase.GetByCatID(c, catID, userID)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", data)
}

// Create godoc
// @description Create new treatment
// @tags  treatment
// @security ApiKeyAuth
// @id TreatmentCreateHandler
// @accept json
// @produce json
// @param cat_id path int true "cat id"
// @param createTreatment body payload.CreateTreatment true "create treatment"
// @Router /api/treatment/{cat_id} [post]
func (t *TreatmentHandler) Create(c *gin.Context) {
	var req payload.CreateTreatment
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, err)
		return
	}

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	catID, err := strconv.ParseInt(c.Param("cat_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	data, err := mapping.Mapping[domain.Treatment](&req)
	if err != nil {
		response.NewError(c, err)
		return
	}

	err = t.treatmentUsecase.Create(c, catID, userID, data)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusCreated, "success", nil)
}

// Update godoc
// @description Update treatment by ID
// @tags  treatment
// @security ApiKeyAuth
// @id TreatmentUpdateHandler
// @accept json
// @produce json
// @param cat_id path int true "cat id"
// @param id path int true "id of treatment"
// @param updateTreatment body payload.UpdateTreatment true "update treatment"
// @Router /api/treatment/{cat_id}/{id} [put]
func (t *TreatmentHandler) Update(c *gin.Context) {
	var req payload.UpdateTreatment
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, err)
		return
	}

	catID, err := strconv.ParseInt(c.Param("cat_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	data, err := mapping.Mapping[domain.Treatment](&req)
	if err != nil {
		response.NewError(c, err)
		return
	}

	err = t.treatmentUsecase.Update(c, id, catID, data)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", nil)
}

// Delete godoc
// @description Delete treatment by ID
// @tags  treatment
// @security ApiKeyAuth
// @id TreatmentDeleteHandler
// @accept json
// @produce json
// @param cat_id path int true "cat id"
// @param id path int true "id of treatment"
// @Router /api/treatment/{cat_id}/{id} [delete]
func (t *TreatmentHandler) Delete(c *gin.Context) {
	catID, err := strconv.ParseInt(c.Param("cat_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	err = t.treatmentUsecase.Delete(c, id, catID)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", nil)
}
