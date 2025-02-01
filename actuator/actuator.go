package actuator

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	
	"github.com/dennesshen/photon-core-starter/bean"
	"github.com/dennesshen/photon-core-starter/core"
	"github.com/dennesshen/photon-core-starter/log"
	"github.com/gofiber/fiber/v2"
)

func init() {
	bean.Autowire(&component)
	core.RegisterProjectInit(InitVersion)
}

var component struct {
	Server *fiber.App `autowired:"true" `
}

var version = []byte("ok")

func InitVersion(ctx context.Context) error {
	versionStr, err := os.ReadFile("version.txt")
	if err != nil {
		log.Logger().Warn(ctx, "read version.txt : ", "error", err)
	} else {
		version = versionStr
	}
	return nil
}

func Start(ctx context.Context) error {
	if !config.Actuator.Enabled {
		return nil
	}
	basePath := config.Actuator.BasePath
	contextPath := config.Actuator.ContextPath
	if contextPath == "" {
		contextPath = "/actuator"
	}
	groupPath := basePath + contextPath
	info := fmt.Sprintf("actuator enabled, group path: %s", groupPath)
	slog.InfoContext(ctx, info)
	actuator := component.Server.Group(groupPath)
	actuator.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Send(version)
	})
	return nil
}
