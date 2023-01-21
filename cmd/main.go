package cmd

import (
	"context"
	"log"
	"net/http" //nolint:nolintlint,gofumpt

	"github.com/labstack/echo/v4"

	"prancerTest/internal/feature/agent/usecase" //nolint:nolintlint,gofumpt
	prancerUsecase "prancerTest/internal/feature/prancer/usecase"
	"prancerTest/internal/pancer/entity" //nolint:nolintlint,gofumpt
)

func init() {
	for i := 0; i < entity.AgentCount; i++ {
		agent := usecase.New(i)
		entity.Agents.Store(agent.GetName(context.TODO()), agent)
	}
}

func RunServer() {
	server := prancerUsecase.New(1)

	echoEngine := echo.New()
	echoEngine.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello") //nolint:wrapcheck
	})

	echoEngine.POST("/move", func(c echo.Context) error {
		var move entity.Transfer
		if err := c.Bind(&move); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error()) //nolint:wrapcheck
		}

		go func() {
			err := server.Calculate(context.TODO(), move)
			if err != nil {
				log.Println(err.Error())
			}
		}()

		return c.String(http.StatusOK, "hello") //nolint:wrapcheck
	})

	echoEngine.Logger.Fatal(echoEngine.Start(":18080"))
}
