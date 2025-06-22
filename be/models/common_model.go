package models

import "github.com/gofiber/fiber/v2"

type response struct {
	Code string      `json:"code"`
	Desc string      `json:"desc"`
	Data interface{} `json:"data"`
}

func NewResponse(code, desc string, data interface{}) *response {
	return &response{
		Code: code,
		Desc: desc,
		Data: data,
	}
}

func (r *response) SendResponse(ctx *fiber.Ctx, httpStatus int) error {
	return ctx.Status(httpStatus).JSON(r)
}