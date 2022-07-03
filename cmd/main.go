package main

import (
	"github.com/Fevzik/modiniter"
	simple_go "github.com/Fevzik/simple-go"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := modiniter.InitModule(simple_go.New(), modiniter.ModuleConfig{
		InitMode:      modiniter.InitModeStorelessService,
		RouterInitMode: modiniter.HttpModeFull,
		Server:         app,
		Store:          nil,
		GRPCPort:       nil,
		AMQPDsn:        nil,
		Etcd:           nil,
		ELK:            nil,
	})
	if err != nil {
		panic(err)
	}


	app.Listen(":3000")
}
