package handler

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kongsakchai/catopia-backend/api/response"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

type FileHandler struct {
}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

func (h *FileHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	extension := filepath.Ext(file.Filename)
	if extension != ".jpg" && extension != ".jpeg" && extension != ".png" {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", nil))
		return
	}

	newFileName := uuid.New().String() + extension

	if err := c.SaveUploadedFile(file, "uploads/"+newFileName); err != nil {
		response.NewErrorResponse(c, errs.New(errs.ErrBadRequest, "Bad Request", err))
		return
	}

	response.NewResponse(c, 200, "success", gin.H{"file_name": newFileName})
}
