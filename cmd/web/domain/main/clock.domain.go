package mainDomain

import (
	"net/http"

	"github.com/scrivx/go-api-web/internal/models"
	"github.com/scrivx/go-api-web/internal/utils"
)

func ClockFragment(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.CLOCK)
	utils.ParseTemplateFiles(w, "clock", utils.EmplyStruct, utils.EmplyFuncMap, "iu/html/pages/main/clock.html")
}