package simple_go

import (
	"github.com/Fevzik/modiniter"
	"github.com/Fevzik/simple-go/internal/constants"
	"github.com/Fevzik/simple-go/internal/handler"
	"github.com/Fevzik/simple-go/internal/repository"
	"github.com/Fevzik/simple-go/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Module struct {
	config *modiniter.ModuleConfig
}

func New() interface{} {
	conf := new(modiniter.ModuleConfig)
	return &Module{config: conf}
}


func (m *Module) GetModuleLabel() string {
	return constants.ModuleLabel
}

func (m *Module) GetModuleCode() string {
	return constants.ModuleCode
}

func (m *Module) GetConfig() *modiniter.ModuleConfig {
	return m.config
}

func (m *Module) ImportRoutes(app *fiber.App, mode string) {
	r := repository.NewRepository(m.config.Store)
	s := service.NewService(r)
	h := handler.NewHandler(s)
	h.InitRoutes(app, m.config.RouterInitMode)
}