package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kongsakchai/catopia-backend/src/helper"
)

func Create[Req any](c *gin.Context, service func(*gin.Context, *Req) error) {
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, helper.BadRequestResponse(err))
		return
	}

	if err := service(c, &req); err != nil {
		c.JSON(500, helper.InternalServerErrorResponse(err))
		return
	}

	c.JSON(201, helper.GenerateResponse(true, nil, "success"))
}

func GetByBody[Req any, Res any](c *gin.Context, service func(*gin.Context, *Req) (*Res, error)) {
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, helper.BadRequestResponse(err))
		return
	}

	data, err := service(c, &req)

	if err != nil {
		c.JSON(500, helper.InternalServerErrorResponse(err))
		return
	}

	c.JSON(201, helper.GenerateResponse(true, data, "success"))
}
