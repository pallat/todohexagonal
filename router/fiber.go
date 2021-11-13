package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pallat/todoapi/todo"
)

type FiberRouter struct {
	*fiber.App
}

func NewFiberRouter() *FiberRouter {
	r := fiber.New()

	r.Use(cors.New())
	r.Use(logger.New())

	return &FiberRouter{r}
}

func (r *FiberRouter) POST(path string, handler func(todo.Context)) {
	r.App.Post(path, func(c *fiber.Ctx) error {
		handler(NewFiberCtx(c))
		return nil
	})
}

type FiberCtx struct {
	*fiber.Ctx
}

func NewFiberCtx(c *fiber.Ctx) *FiberCtx {
	return &FiberCtx{Ctx: c}
}

func (c *FiberCtx) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}
func (c *FiberCtx) JSON(statuscode int, v interface{}) {
	c.Ctx.Status(statuscode).JSON(v)
}
func (c *FiberCtx) TransactionID() string {
	return string(c.Ctx.Request().Header.Peek("TransactionID"))
}
func (c *FiberCtx) Audience() string {
	return c.Ctx.Get("aud")
}
