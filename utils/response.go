package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status       string      `json:"status"`
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Error        string      `json:"error,omitempty"`
}

type ResponsePaginated struct {
	Status       string         `json:"status"`
	ResponseCode int            `json:"response_code"`
	Message      string         `json:"message,omitempty"`
	Data         interface{}    `json:"data,omitempty"`
	Error        string         `json:"error,omitempty"`
	Meta         PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	Page      int    `json:"page" example:"1"`
	Limit     int    `json:"limit" example:"10"`
	Total     int    `json:"total" example:"100"`
	TotalPage int    `json:"total_pages" example:"10"`
	Filter    string `json:"filter" example:"nama=ilham"`
	Sort      string `json:"sort" example:"-id"`
}

func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:       "success",
		ResponseCode: fiber.StatusOK,
		Message:      message,
		Data:         data,
	})
}

func SuccessPagination(c *fiber.Ctx, message string, data interface{}, meta PaginationMeta) error {
	return c.Status(fiber.StatusOK).JSON(ResponsePaginated{
		Status:       "success",
		ResponseCode: fiber.StatusOK,
		Message:      message,
		Data:         data,
		Meta:         meta,
	})
}

func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Status:       "Created",
		ResponseCode: fiber.StatusCreated,
		Message:      message,
		Data:         data,
	})
}

func BadRequest(c *fiber.Ctx, message string, error string) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Status:       "Error: Bad Request",
		Error:        error,
		ResponseCode: fiber.StatusBadRequest,
		Message:      message,
	})
}

func NotFound(c *fiber.Ctx, message string, error string) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Status:       "Error: Not Found",
		Error:        error,
		ResponseCode: fiber.StatusNotFound,
		Message:      message,
	})
}

func NotFoundPagination(c *fiber.Ctx, message string, data interface{}, meta PaginationMeta) error {
	return c.Status(fiber.StatusNotFound).JSON(ResponsePaginated{
		Status:       "Error: Not Found",
		ResponseCode: fiber.StatusNotFound,
		Message:      message,
		Data:         data,
		Meta:         meta,
	})
}

func Unauthorized(c *fiber.Ctx, message string, error string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		Status:       "Error: Not Found",
		Error:        error,
		ResponseCode: fiber.StatusUnauthorized,
		Message:      message,
	})
}

func InternalServerError(c *fiber.Ctx, message string, error string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		Status:       "Error: Internal Server Error",
		Error:        error,
		ResponseCode: fiber.StatusInternalServerError,
		Message:      message,
	})
}


