package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/response"
	"github.com/kongsakchai/catopia-backend/domain"
)

type RecommendHandler struct {
	recommendUsecase domain.RecommendUsecase
}

func NewRecommendHandler(recommendUsecase domain.RecommendUsecase) *RecommendHandler {
	return &RecommendHandler{recommendUsecase}
}

// Get By Cat ID godoc
// @description Get cat by cat ID
// @tags recommend
// @security ApiKeyAuth
// @id RecommendGetByCatIDHandler
// @accept json
// @produce json
// @param id path int true "id of cat"
// @Router /api/recommend/cat/{id} [get]
func (h *RecommendHandler) GetByCatID(c *gin.Context) {
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

	data, err := h.recommendUsecase.RecommendByCat(c, id, userID)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", data)
}

// Get By User godoc
// @description Get cat by user data
// @tags recommend
// @security ApiKeyAuth
// @id RecommendGetByUserHandler
// @accept json
// @produce json
// @Router /api/recommend/cat [get]
func (h *RecommendHandler) GetByUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		response.NewError(c, err)
		return
	}

	data, err := h.recommendUsecase.RecommendByUser(c, userID)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", data)
}
