package mainRoute

import (
	mainDomain "github.com/scrivx/go-api-web/cmd/web/domain/main"
	"github.com/scrivx/go-api-web/cmd/web/routes"
	"github.com/scrivx/go-api-web/internal/models"
	"github.com/scrivx/go-api-web/internal/utils"
)

func ClockRender() {
	duration := mainDomain.GetDuration()
	clock := mainDomain.GetClockInstance()

	utils.SetDuration(duration)
	go utils.StartCountdown(clock, duration)
	// usar "GET" sin espacios puede fallar, pero usar "GET " con espacios funciona
	routes.GetMuxInstance().HandleFunc("GET " + models.RoutesInstance.CLOCK, mainDomain.ClockFragment)
	routes.GetMuxInstance().HandleFunc("POST " + models.RoutesInstance.RESTART_CYCLE, mainDomain.RestartCycle)
}