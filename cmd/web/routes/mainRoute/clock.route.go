package mainRoute

import (
	"github.com/scrivx/go-api-web/cmd/web/routes"
	"github.com/scrivx/go-api-web/internal/models"
)

func ClockRender() {
	routes.GetMuxInstance().HandleFunc("GET" + models.RoutesInstance.CLOCK, clockDomain.ClockView)
}