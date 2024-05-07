package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/response"
	"github.com/kongsakchai/catopia-backend/domain"
)

type RecommendHandler struct {
	catUsecase domain.CatUsecase
}

func NewRecommendHandler(catUsecase domain.CatUsecase) *RecommendHandler {
	return &RecommendHandler{catUsecase}
}

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

	data, err := h.catUsecase.GetCluster(c, id, userID)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", data)
}

func (h *RecommendHandler) GetRandom(c *gin.Context) {
	data, err := h.catUsecase.GetRandom(c)
	if err != nil {
		response.NewError(c, err)
		return
	}

	response.New(c, http.StatusOK, "success", data)
}
