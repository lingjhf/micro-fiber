package microfiber

import "github.com/gofiber/fiber/v2"

var (
	ErrUnknown      = NewError("unknow", "未知错误")
	ErrSucceeded    = NewError("succeeded", "成功")
	ErrBodyParser   = NewError("err_body_parser", "请求体解析错误")
	ErrQueryParser  = NewError("err_query_parser", "查询参数解析错误")
	ErrParameter    = NewError("err_parameter", "参数错误")
	ErrDatabase     = NewError("err_database", "数据库操作错误")
	ErrUnauthorized = NewError("err_unauthorized", "未授权")
	ErrForbidden    = NewError("err_forbidden", "无权限")
)

type ErrorMessage struct {
	Code    string
	Message any
}

func NewError(code string, message any) *ErrorMessage {
	return &ErrorMessage{Code: code, Message: message}
}

func Message(key string, value string) map[string]string {
	return map[string]string{key: value}
}

type JSONResponse struct {
	Code    string `json:"code"`
	Message any    `json:"message"`
	Data    any    `json:"data"`
}

func Succeeded(c *fiber.Ctx) error {
	return c.JSON(&JSONResponse{
		Code:    ErrSucceeded.Code,
		Message: ErrSucceeded.Message,
	})
}

func SucceededWithData(c *fiber.Ctx, data any) error {
	return c.JSON(&JSONResponse{
		Code:    ErrSucceeded.Code,
		Message: ErrSucceeded.Message,
		Data:    data,
	})
}

func BodyParserError(c *fiber.Ctx, message any) error {
	return c.Status(fiber.StatusBadRequest).JSON(&JSONResponse{
		Code:    ErrBodyParser.Code,
		Message: message,
	})
}

func QueryParserError(c *fiber.Ctx, message any) error {
	return c.Status(fiber.StatusBadRequest).JSON(&JSONResponse{
		Code:    ErrQueryParser.Code,
		Message: message,
	})
}

func ParameterError(c *fiber.Ctx, message any) error {
	return c.Status(fiber.StatusBadRequest).JSON(&JSONResponse{
		Code:    ErrParameter.Code,
		Message: message,
	})
}

func UnauthorizedError(c *fiber.Ctx, message ...any) error {
	m := ErrUnauthorized.Message
	if len(message) > 0 {
		m = message[0]
	}
	return c.Status(fiber.StatusUnauthorized).JSON(&JSONResponse{
		Code:    ErrUnauthorized.Code,
		Message: m,
	})
}

func ForbiddenError(c *fiber.Ctx, message ...any) error {
	m := ErrForbidden.Message
	if len(message) > 0 {
		m = message[0]
	}
	return c.Status(fiber.StatusForbidden).JSON(&JSONResponse{
		Code:    ErrForbidden.Code,
		Message: m,
	})
}

func DatabaseError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(&JSONResponse{
		Code:    ErrDatabase.Code,
		Message: ErrDatabase.Message,
	})
}

func UnknownError(c *fiber.Ctx, message ...any) error {
	m := ErrUnknown.Message
	if len(message) > 0 {
		m = message[0]
	}
	return c.Status(fiber.StatusInternalServerError).JSON(&JSONResponse{
		Code:    ErrUnknown.Code,
		Message: m,
	})
}
