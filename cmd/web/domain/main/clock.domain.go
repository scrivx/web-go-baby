package mainDomain

import (
	"net/http"

	"github.com/scrivx/go-api-web/internal/models"
	"github.com/scrivx/go-api-web/internal/utils"
	"github.com/scrivx/go-api-web/ui"
)

var duration = 14400

var clockInstance = utils.NewClock()

func GetClockInstance() *utils.Clock {
	return clockInstance
}

func GetDuration() int {
	return duration
}

func ClockFragment(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.CLOCK)
	utils.ParseTemplateFiles(w, "clock", clockInstance, utils.EmplyFuncMap, ui.Content, "html/pages/main/clock.html")
}

func RestartCycle(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.RESTART_CYCLE)
	
	select {
	case <- clockInstance.Stop: // si el canal ya esta cerrado, no hace nada
	default: // si el canal no esta cerrado
		// guardamos el valor del log en un base de datos
		utils.StopCountdown(clockInstance)
		clockInstance.CountDown = "04:00:00"
		utils.SetDuration(duration)
		go utils.StartCountdown(clockInstance, duration)
		
		w.Write([]byte("Cycle restarted"))
	}
}