package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/api/payload"
	"github.com/kongsakchai/catopia-backend/api/response"
	"github.com/kongsakchai/catopia-backend/domain"
	errs "github.com/kongsakchai/catopia-backend/domain/error"
)

type OTPHandler struct {
	otpUsecase domain.OTPUsecase
}

func NewOTPHandler(otpUsecase domain.OTPUsecase) *OTPHandler {
	return &OTPHandler{otpUsecase}
}

// VerifyOTP godoc
// @description Verify OTP
// @tags Forgot Password
// @id VerifyOTPHandler
// @accept json
// @produce json
// @param code body payload.VerifyOTP true "code"
// @Router /api/otp/verify [post]
func (h *OTPHandler) VerifyOTP(c *gin.Context) {
	var req payload.VerifyOTP
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewError(c, err)
		return
	}

	otp, notExpire, err := h.otpUsecase.GetOTP(c, req.Code)
	if err != nil {
		response.NewError(c, err)
		return
	}

	if !notExpire {
		response.NewError(c, errs.NewError(errs.ErrNotFound, fmt.Errorf("otp expired")))
		h.otpUsecase.Delete(c, req.Code)
		return
	}

	if otp.OTP != req.OTP {
		response.NewError(c, errs.NewError(errs.ErrNotFound, fmt.Errorf("otp not match")))
		return
	}

	response.New(c, http.StatusOK, "success", nil)
}
