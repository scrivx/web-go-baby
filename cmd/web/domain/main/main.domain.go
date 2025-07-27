package mainDomain

import (
	"net/http"

	"github.com/scrivx/go-api-web/internal/models"
	"github.com/scrivx/go-api-web/internal/utils"
)

func MainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.MAIN)
	files := []string{
		"ui/html/base.html",
		"ui/html/pages/main/main.html",
	}
	
	utils.ParseTemplateFiles(w, "base", utils.EmplyStruct, utils.EmplyFuncMap, files...)
}

